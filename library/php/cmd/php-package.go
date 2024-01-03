/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"projectista/library"
	"projectista/library/php"
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

		project := library.Project{Name: args[0]}
		project.RootFolder, _ = cmd.Flags().GetString("folder")

		println("Scaffolding the project " + project.Name + " in the folder " + project.RootFolder)

		php.Build(project)
	},
}
