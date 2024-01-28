package library

import (
	"projectista/parser"
)

/*
Scaffolding logic
*/

func Scaffold(sourceDirectory string, outDirectory string, parameters map[string]string, excludedFiles []string) bool {

	rHelper := parser.New(sourceDirectory, outDirectory)
	rHelper.Walk(parameters, excludedFiles)

	return true
}
