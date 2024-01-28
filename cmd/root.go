/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>

This file contains the main root command.
Commands for specific project / language are defined in the
library folder
*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"projectista/library/laravel"
	"projectista/library/php"
)

var version = "1.0.4"

var rootCmd = &cobra.Command{
	Use:   "pi",
	Version: version,
	Short: "Projectista is a simple command line tool to scaffold projects",
	Long: `Projectista is a simple command line tool to scaffold projects.
For example:

pi php package mypackage
pi php package mypackage --folder=mypackage
`,
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// The folder in which the project will be created
	// If not provided, the project will be created in the current folder
	rootCmd.PersistentFlags().String("folder", ".", "The folder in which the project will be created")

	// Set the version template
	rootCmd.SetVersionTemplate("Projectista version {{.Version}}\n")

	// Register sub-commands
	rootCmd.AddCommand(php.Cmd)
	rootCmd.AddCommand(laravel.Cmd)
}
