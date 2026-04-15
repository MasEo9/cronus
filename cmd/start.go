/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"cronus/internal/models"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time for project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		newSession := models.Session{
			Date: time.Now().Format("2006.01.02"),
		}
		newSession.TimeStart = time.Now()

		err := appDB.InsertSession(projectName, newSession.Date, newSession.TimeStart)
		if err != nil {
			fmt.Println("Unable to insert new session %v\n", err)
			return 
		}
		fmt.Printf("Started tracking project %s at %s\n", projectName, newSession.TimeStart.Format(time.Kitchen))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to start")
}
