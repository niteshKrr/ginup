package cmd

import (
	"fmt"

	"github.com/niteshKrr/ginup/internal"
	"github.com/niteshKrr/ginup/internal/scaffolder"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Gin backend project",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Error: Project name required")
			return
		}

		projectName := args[0]

		projectPath, err := internal.BasicSetup(projectName)
		if err != nil {
			fmt.Println(err)
			return
		}
		scaffolder.CreateProject(projectPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
