/*
Copyright © 2026 MasEo9 <ajmason09@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"cronus/internal/models"
	"github.com/spf13/cobra"
)

// input model
type projectAddModel struct {
	projectName string
	choices     []string
	cursor      int
	choice      string
}

// init method required by tea
func (m projectAddModel) Init() tea.Cmd {
	return nil
}

func (m projectAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			m.choice = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m projectAddModel) View() tea.View {
	s := fmt.Sprintf("Project [%s] does not exist. What would you like to do?\n", m.projectName)

	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "Press q to quit.\n"
	return tea.NewView(s)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time for project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")

		p, err := appDB.SearchProjects(projectName)

		if p.ProjectName != projectName {
			initialModel := projectAddModel{
				projectName: projectName,
				choices:     []string{"Add Project", "Exit"},
			}

			pTUI := tea.NewProgram(initialModel)
			finalModel, err := pTUI.Run()
			if err != nil {
				fmt.Printf("Error rendering UI: %v\n", err)
				os.Exit(1)
			}

			m := finalModel.(projectAddModel)

			if m.choice == "Exit" || m.choice == "" {
				fmt.Println("Exiting without starting session.")
				return
			}

			if m.choice == "Add Project" {
				// Insert the project and grab the ID
				newID, err := appDB.InsertProject(projectName)
				if err != nil {
					fmt.Printf("Failed to create project: %v\n", err)
					return
				}
				// Assign the new ID 
				p.ID = newID
				p.ProjectName = projectName
				fmt.Printf("Created new project: %s\n", projectName)
			}

		} else if err != nil {
			fmt.Printf("Database error: %v\n", err)
			return
		}

		newSession := models.Session{
			Date: time.Now().Format("2006.01.02"),
		}
		newSession.TimeStart = time.Now()

		err = appDB.InsertSession(p.ID, newSession.Date, newSession.TimeStart)
		if err != nil {
			fmt.Printf("Unable to insert new session %v\n", err)
			return
		}
		fmt.Printf("Started tracking project %s at %s\n", projectName, newSession.TimeStart.Format(time.Kitchen))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("project", "p", "Unnamed", "Name of the project to start")
}
