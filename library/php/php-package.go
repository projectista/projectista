/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"bytes"
	"github.com/spf13/cobra"
	"os"
	"projectista/reader"
	"projectista/stubs"
	"projectista/writer"
	"text/template"
)

/*
Cobra command
*/

var PackageCmd = &cobra.Command{
	Use:   "package [projectName]",
	Args:  cobra.ExactArgs(1),
	Short: "Create a PHP Package",
	Long: `Create a PHP Package.
This command will create a PHP Package with
- A
- B
- C`,
	Run: func(cmd *cobra.Command, args []string) {

		var projectName = args[0]
		var outDirectory, _ = cmd.Flags().GetString("folder")

		parameters := make(map[string]string)
		parameters["Name"] = projectName

		println("Scaffolding the project " + projectName + " in the folder " + outDirectory)

		scaffold(outDirectory, parameters)
	},
}

/*
Scaffolding logic
*/

func scaffold(outDirectory string, parameters map[string]string) bool {

	var sourceDirectory = "stubs/php/package/"
	var files = []string{
		"composer.json",
	}

	_ = writer.New(outDirectory)
	_ = reader.New(sourceDirectory)

	for _, file := range files {

		fileContent, err := stubs.FS.ReadFile(sourceDirectory + file)
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

		err = os.WriteFile(outDirectory+string(os.PathSeparator)+file, buffer.Bytes(), 0644)
		if err != nil {
			return false
		}
	}

	return true
}
