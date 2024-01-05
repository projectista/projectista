/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"github.com/spf13/cobra"
	"projectista/parser"
	"projectista/printer"
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
		parameters["ProjectName"] = projectName
		parameters["VendorName"] = "Vendor"

		printer.Header("Scaffolding php-package: \nProject: " + projectName + "\n" + "Directory: " + outDirectory)

		scaffold(outDirectory, parameters)
	},
}

/*
Scaffolding logic
*/

func scaffold(outDirectory string, parameters map[string]string) bool {

	var sourceDirectory = "stubs/php/php-package"

	var excludedFiles = []string{
		"/build.sh",
	}

	rHelper := parser.New(sourceDirectory, outDirectory)
	rHelper.Walk(parameters, excludedFiles)

	return true
}
