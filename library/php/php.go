package php

import (
	"bytes"
	"os"
	"projectista/library"
	"projectista/stubs"
	"text/template"
)

var rootDir = "stubs/php/package/"
var files = []string{
	"composer.json",
}

func Build(project library.Project) bool {

	for _, file := range files {

		fileContent, err := stubs.FS.ReadFile(rootDir + file)
		if err != nil {
			panic(err)
		}

		fileTemplate, err := template.New("tpl").Parse(string(fileContent))
		if err != nil {
			panic(err)
		}

		buffer := new(bytes.Buffer)
		err = fileTemplate.Execute(buffer, project)
		if err != nil {
			panic(err)
		}

		if _, err := os.Stat(project.RootFolder); os.IsNotExist(err) {
			println("Folder", project.RootFolder, "does not exists")
			os.Exit(1)
		}

		err = os.WriteFile(project.RootFolder+string(os.PathSeparator)+file, buffer.Bytes(), 0644)
		if err != nil {
			return false
		}
	}

	return true
}
