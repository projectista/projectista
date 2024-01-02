package php

import t "projectionista/templates"

var rootDir = "templates/php/package/"
var files = []string{
	"composer.json",
}

func Build() bool {

	for _, file := range files {

		data, _ := t.FS.ReadFile(rootDir + file)
		println(string(data))
	}

	return true
}
