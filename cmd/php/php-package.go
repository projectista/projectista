/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"github.com/spf13/cobra"
	"projectista/php"
)

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
		var projectFolder, _ = cmd.Flags().GetString("folder")

		println("Scaffolding the project " + projectName + " in the folder " + projectFolder)

		php.Build()
	},
}
