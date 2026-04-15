/*
Copyright © 2026 MasEo9 <ajmason09@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		s, err := appDB.GetActiveSession(projectName)
		if err != nil {
			fmt.Println("Unable to get active session %w", err)
			return
		}

		s.TimeEnd = time.Now()
		elapsedDuration := s.CalculateElapsed()

		fmt.Printf("Current Session time elapsed: %s, Session started at: %s\n", elapsedDuration.Round(time.Second), s.TimeStart.Format(time.Kitchen))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to start")
}
