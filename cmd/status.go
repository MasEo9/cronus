/*
Copyright © 2026 MasEo9 <ajmason09@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"
	"errors"
	"database/sql"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the elapsed time of an actively running project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		s, err := appDB.GetActiveSession(projectName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				fmt.Printf("There is no active session running for project '%s'.\n", projectName)
				return
			}
			fmt.Printf("Database error while checking session: %v\n", err)
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
