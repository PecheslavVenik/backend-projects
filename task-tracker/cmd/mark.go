package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker/internal/task"
)

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.Done)
		},
	}
	return cmd
}

func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as in-progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.InProgress)
		},
	}
	return cmd
}

func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.ToDo)
		},
	}
	return cmd
}

func RunUpdateStatusCmd(args []string, status task.Status) error {
	if len(args) == 0 {
		return fmt.Errorf("task ID is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return task.UpdateTaskStatus(id, status)
}
