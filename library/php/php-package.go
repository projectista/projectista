/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"github.com/spf13/cobra"
	"projectista/reader"
	"projectista/writer"
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

	var sourceDirectory = "stubs/php/php-package/"
	var files = []string{
		"composer.json",
	}

	wHelper := writer.New(outDirectory)
	rHelper := reader.New(sourceDirectory)

	for _, file := range files {
		pathToFile := wHelper.Write(file, rHelper.Parse(file, parameters))

		println("Written", pathToFile)
	}

	return true
}
