/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package laravel

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
	Short: "Create a new Laravel Package",
	Long: `
Create a new Laravel Package.

This command will create a Laravel Package that includes:

- Pest testing framework
- PHPStan static analysis
- Pint integrtion

The stubs are available on GitHub: https://github.com/projectista/laravel-package
`,
	Run: func(cmd *cobra.Command, args []string) {

		sourceDirectory := "stubs/laravel/laravel-package"
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

		printer.Header("Scaffolding laravel application: \nProject: " + parameters["ProjectName"] + "\n" + "Directory: " + outDirectory)

		library.Scaffold(sourceDirectory, outDirectory, parameters, excludedFiles)
	},
}

func init() {

	PackageCmd.Flags().StringP("description", "d", "A wonderful Projectista package", "The description for the project")
	PackageCmd.Flags().StringP("vendor", "v", "Projectista", "The name of the vendor owning the project")
	PackageCmd.Flags().StringP("author", "a", "Projectista User", "The name of the author of the package")
	PackageCmd.Flags().StringP("email", "e", "me@example.com", "The email of the author of the package")
}
