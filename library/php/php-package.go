/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"github.com/spf13/cobra"
	"projectista/library"
	"projectista/printer"
)

/*
Cobra command
*/

var PackageCmd = &cobra.Command{
	Use:   "package [projectName]",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new PHP Package",
	Long: `
Create a new PHP Package.

This command will create a PHP Package that includes:

- Pest testing framwork
- PHPStan static analysis
- Pint integrtion

The stubs are available on GitHub: https://github.com/projectista/php-package
`,
	Run: func(cmd *cobra.Command, args []string) {

		sourceDirectory := "stubs/php/php-package"
		outDirectory, _ := cmd.Flags().GetString("folder")

		parameters := make(map[string]string)
		parameters["ProjectName"] = args[0]
		parameters["ProjectDescription"], _ = cmd.LocalFlags().GetString("description")
		parameters["VendorName"], _ = cmd.LocalFlags().GetString("vendor")
		parameters["AuthorName"], _ = cmd.LocalFlags().GetString("author")
		parameters["AuthorEmail"], _ = cmd.LocalFlags().GetString("email")

		excludedFiles := []string{
			"/build.sh",
		}

		printer.Header("Scaffolding php-package: \nProject: " + parameters["ProjectName"] + "\n" + "Directory: " + outDirectory)

		library.Scaffold(sourceDirectory, outDirectory, parameters, excludedFiles)
	},
}

func init() {

	PackageCmd.Flags().StringP("description", "d", "A wonderful Projectista package", "The description for the project")
	PackageCmd.Flags().StringP("vendor", "v", "Projectista", "The name of the vendor owning the project")
	PackageCmd.Flags().StringP("author", "a", "Projectista User", "The name of the author of the package")
	PackageCmd.Flags().StringP("email", "e", "me@example.com", "The email of the author of the package")
}
