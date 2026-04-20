/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, _ := appDB.ListProjects()
		fmt.Print("ID | Project \n")
		for _, p := range projects {
			fmt.Printf("%v, %s \n", p.ID, p.ProjectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
