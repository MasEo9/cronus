/*
Copyright © 2026 MasEo9 <ajmason09@gmail.com>
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

		sessions, err := appDB.GetActiveSession()
		if err != nil {
			fmt.Printf("Database error while checking sessions: %v\n", err)
			return
		}

		if len(sessions) == 0 {
			fmt.Println("There are no active sessions currently running.")
			return
		}

		for _, s := range sessions {
			// project name is passed in and the project exists and is found (data = flag)
			if s.ProjectName != projectName {
				continue
			}

			s.TimeEnd = time.Now()
			elapsedDuration := s.CalculateElapsed()

			err = appDB.StopSession(s.TimeEnd, elapsedDuration, float32(elapsedDuration.Hours()), s.ID)
			if err != nil {
				fmt.Printf("Unable to end session %v\n", err)
				return
			}

			fmt.Printf("Session ended at: %s, elapsed time: %s\n", s.TimeEnd.Format(time.Kitchen), elapsedDuration.Round(time.Second))
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to stop")
}
