/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"github.com/spf13/cobra"
)

var PhpCmd = &cobra.Command{
	Use:   "php",
	Short: "Scaffold a PHP project",
	Long: `Scaffold a PHP project.
This command will scaffold a PHP project.
You need to specify which kind of project you want to
create.

Example:

projectionista php package my-package-name`,
}

func init() {
	PhpCmd.AddCommand(PackageCmd)
}
