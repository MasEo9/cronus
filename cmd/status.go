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
	Short: "Check the elapsed time of an actively running project",
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

		foundTargetProject := false

		for _, s := range sessions {
			// project name is passed in and the project exists and is found (data = flag)
			if projectName != "Unnamed" && s.ProjectName != projectName {
				continue
			}


			foundTargetProject = true
			s.TimeEnd = time.Now()
			elapsedDuration := s.CalculateElapsed()

			fmt.Printf("[%s] time elapsed: %s, session started at: %s\n",
				s.ProjectName,
				elapsedDuration.Round(time.Second),
				s.TimeStart.Format(time.Kitchen))
		}

		// no found project and no flag
		if !foundTargetProject && projectName != "Unnamed" {
			fmt.Printf("There is no active session running for project '%s'.\n", projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project")
}
