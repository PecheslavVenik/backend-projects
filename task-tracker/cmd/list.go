package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status

    Example:
    ./task-tracker list todo
    ./task-tracker list in-progress
    ./task-tracker list done
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				status := task.Status(args[0])
				return task.ListTasks(status)
			}

			return task.ListTasks("all")
		},
	}
	return cmd
}
