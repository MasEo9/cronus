/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		err := appDB.InsertProject(projectName)
		if err != nil {
			fmt.Printf("Unable to insert new project %v\n", err)
			return
		}
		fmt.Printf("[%s] project added\n", projectName)
	},
}

func init() {
	projectCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to start")
}
