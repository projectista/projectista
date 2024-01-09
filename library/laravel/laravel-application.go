/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package laravel

import (
	"github.com/spf13/cobra"
	"projectista/parser"
	"projectista/printer"
)

/*
Cobra command
*/

var ApplicationCmd = &cobra.Command{
	Use:   "application [projectName]",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new Laravel Application",
	Long: `
Create a new Laravel Application.

This command will create a Laravel Application that includes:

- Laravel 10
- Pest testing framework
- PHPStan static analysis
- Pint integrtion

The stubs are available on GitHub: https://github.com/projectista/laravel
`,
	Run: func(cmd *cobra.Command, args []string) {

		outDirectory, _ := cmd.Flags().GetString("folder")

		parameters := make(map[string]string)
		parameters["ProjectName"] = args[0]
		parameters["ProjectDescription"], _ = cmd.LocalFlags().GetString("description")
		parameters["VendorName"], _ = cmd.LocalFlags().GetString("vendor")
		parameters["AuthorName"], _ = cmd.LocalFlags().GetString("author")
		parameters["AuthorEmail"], _ = cmd.LocalFlags().GetString("email")

		printer.Header("Scaffolding laravel application: \nProject: " + parameters["ProjectName"] + "\n" + "Directory: " + outDirectory)

		scaffold(outDirectory, parameters)
	},
}

func init() {

	ApplicationCmd.Flags().StringP("description", "d", "A wonderful Projectista package", "The description for the project")
	ApplicationCmd.Flags().StringP("vendor", "v", "Projectista", "The name of the vendor owning the project")
	ApplicationCmd.Flags().StringP("author", "a", "Projectista User", "The name of the author of the package")
	ApplicationCmd.Flags().StringP("email", "e", "me@example.com", "The email of the author of the package")
}

/*
Scaffolding logic
*/

func scaffold(outDirectory string, parameters map[string]string) bool {

	var sourceDirectory = "stubs/laravel/laravel"

	var excludedFiles = []string{
		"/build.sh",
	}

	rHelper := parser.New(sourceDirectory, outDirectory)
	rHelper.Walk(parameters, excludedFiles)

	return true
}
