package cmd

import (
	"errors"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		Long: `Add a task to the task list. You can provide a description for the task
	Example:
	./task-tracker add 'new task description'`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("task description is required")
			}

			description := args[0]
			return task.AddTask(description)
		},
	}
	return cmd
}
