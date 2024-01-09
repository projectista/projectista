/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package laravel

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "laravel",
	Short: "Scaffold a Laravel project",
	Long: `Scaffold a Laravel project.
This command will scaffold a Laravel project.
You need to specify which kind of project you want to
create.

Example:

projectista laravel dist my-package-name`,
}

func init() {

	Cmd.AddCommand(ApplicationCmd)
}
