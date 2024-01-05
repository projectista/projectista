package parser

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/fs"
	"os"
	"projectista/stubs"
	"slices"
	"strings"
	"text/template"
)

// Parser is the Helper used to read / write and template files
type Parser struct {
	sourceDirectory string
	outDirectory    string
}

// New creates a new instance of the Parser helper
func New(sourceDirectory string, outDirectory string) *Parser {

	t := &Parser{sourceDirectory: sourceDirectory, outDirectory: outDirectory}

	return t
}

// SourceDirectory returns the directory from which the files will be read.
func (t *Parser) SourceDirectory() string {

	return t.sourceDirectory
}

// OutDirectory returns the directry in which the files will be written
func (t *Parser) OutDirectory() string {

	return t.outDirectory
}

// Walk will walk recursively the sourceDirectory,
// foreach file it will template it and write in the outDirectory
func (t *Parser) Walk(parameters map[string]string, excludedFiles []string) {

	err := fs.WalkDir(stubs.FS, t.sourceDirectory, func(path string, dirEntry fs.DirEntry, err error) error {

		cleanPath := strings.Replace(path, t.sourceDirectory, "", 1)

		if dirEntry.IsDir() {
			if path != t.sourceDirectory {
				t.ensureDirectoryExists(cleanPath)
			}
		} else {
			// Parse the current file only if it's not in the excludedFiles slice
			if !slices.Contains(excludedFiles, cleanPath) {
				pathToFile := t.write(cleanPath, t.parse(cleanPath, parameters))
				println("Written", pathToFile)
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

// parse will read a file and template it, using the provided parameters,
// returning an array of bytes
func (t *Parser) parse(file string, parameters map[string]string) []byte {

	titleCaser := cases.Title(language.Und)

	funcMap := template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToUpper":      strings.ToUpper,
		"Title":        titleCaser.String,
		"NoWS":         noWhiteSpaces,
		"Alphanumeric": alphanumeric,
	}

	fileContent, err := stubs.FS.ReadFile(t.sourceDirectory + file)
	if err != nil {
		panic(err)
	}

	fileTemplate, err := template.New("tpl").Funcs(funcMap).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	err = fileTemplate.Execute(buffer, parameters)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

// ensureDirectoryExists will create the provided directory in the outDirectory tree,
// it doesn't exists yet
func (t *Parser) ensureDirectoryExists(directory string) {

	directoryPath := t.outDirectory + directory

	if _, statErr := os.Stat(directoryPath); os.IsNotExist(statErr) {
		err := os.Mkdir(directoryPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

}

// write will write the provided bytes into the provided file
func (t *Parser) write(file string, fileBytes []byte) string {

	completePathToFile := t.outDirectory + file

	err := os.WriteFile(
		completePathToFile, fileBytes, os.ModePerm,
	)
	if err != nil {
		panic(err)
	}

	return completePathToFile
}
