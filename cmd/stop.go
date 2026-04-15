/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop current project timer for the passed in project name",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		s, err := appDB.GetActiveSession(projectName)
		if err != nil {
			fmt.Println("Unable to get active session %w", err)
			return
		}
		fmt.Println("Current active session id", s.ID)

		s.TimeEnd = time.Now()
		elapsedDuration := s.CalculateElapsed()

		appDB.StopSession(s.TimeEnd, elapsedDuration, float32(elapsedDuration.Hours()), s.ID)
		if err != nil {
			fmt.Println("Unable to end session %w", err)
			return
		}
		fmt.Printf("Session ended at: %s, elapsed time: %f minutes\n", s.TimeEnd.Format(time.Kitchen), elapsedDuration.Minutes())
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to stop")
}
