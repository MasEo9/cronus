/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Print out a report of the sessions for the specified time frame",
	Long: `Report out of the cronus database a summary of the sessions as filtered by the 
	passed in flags. Reports can be summarized by project, by day or all time. Weekly or daily summaries 
	can also be generated`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

}
