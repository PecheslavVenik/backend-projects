package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"task-tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}
