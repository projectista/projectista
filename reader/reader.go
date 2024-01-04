package reader

import (
	"bytes"
	"projectista/stubs"
	"text/template"
)

// Reader is an helper to read
// files from the filesystem, and using
// them as templates, using the go templating system
type Reader struct {
	sourceDirectory string
}

// New creates a new instance of the Reader helper
func New(sourceDirectory string) *Reader {
	t := &Reader{sourceDirectory: sourceDirectory}

	return t
}

// SourceDirectory returns the directory from which the files will be read.
func (t *Reader) SourceDirectory() string {
	return t.sourceDirectory
}

// Parse will read a file and template it, using the provided parameters,
// returning an array of bytes
func (t *Reader) Parse(file string, parameters map[string]string) []byte {

	fileContent, err := stubs.FS.ReadFile(t.sourceDirectory + file)
	if err != nil {
		panic(err)
	}

	fileTemplate, err := template.New("tpl").Parse(string(fileContent))
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
