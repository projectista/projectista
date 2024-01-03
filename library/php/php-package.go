/*
Copyright © 2024 Vincenzo Petrucci <vincenzo.petrucci@gmail.com>
*/

package php

import (
	"bytes"
	"github.com/spf13/cobra"
	"os"
	"projectista/project"
	"projectista/stubs"
	"text/template"
)

/*
Cobra command
*/

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
		var rootFolder, _ = cmd.Flags().GetString("folder")

		project := project.New(projectName, rootFolder)

		println("Scaffolding the project " + project.Name() + " in the folder " + project.RootFolder())

		scaffold(project)
	},
}

/*
Scaffolding logic
*/

func scaffold(project *project.Project) bool {

	var rootDir = "stubs/php/package/"
	var files = []string{
		"composer.json",
	}

	if _, err := os.Stat(project.RootFolder()); os.IsNotExist(err) {
		println("Folder", project.RootFolder(), "does not exists")
		os.Exit(1)
	}

	for _, file := range files {

		fileContent, err := stubs.FS.ReadFile(rootDir + file)
		if err != nil {
			panic(err)
		}

		fileTemplate, err := template.New("tpl").Parse(string(fileContent))
		if err != nil {
			panic(err)
		}

		buffer := new(bytes.Buffer)
		err = fileTemplate.Execute(buffer, project)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(project.RootFolder()+string(os.PathSeparator)+file, buffer.Bytes(), 0644)
		if err != nil {
			return false
		}
	}

	return true
}
