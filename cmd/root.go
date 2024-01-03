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
	php "projectista/library/php/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "projectista",
	Short: "A simple command line tool to scaffold projects",
	Long: `A simple command line tool to scaffold projects.
For example:

projectista php package mypackage
projectista php package mypackage --folder=mypackage
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

	// Register sub-commands
	rootCmd.AddCommand(php.Cmd)
}
