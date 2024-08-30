package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Long: `Update a task by providing the task ID and the new status
    Example:
    ./task-tracker update 1 'new description'
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("please provide a task id and the new description")
			}

			taskID := args[0]
			taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
			if err != nil {
				return err
			}

			newDescription := args[1]
			return task.UpdateTaskDescription(taskIDInt, newDescription)
		},
	}

	return cmd
}
