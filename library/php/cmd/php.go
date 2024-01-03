/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "php",
	Short: "Scaffold a PHP project",
	Long: `Scaffold a PHP project.
This command will scaffold a PHP project.
You need to specify which kind of project you want to
create.

Example:

projectista php package my-package-name`,
}

func init() {
	Cmd.AddCommand(PackageCmd)
}
