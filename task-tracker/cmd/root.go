package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
		Long: `Task Tracker is a CLI tool for managing tasks.
		It allows you to create, list, delete, update and mark the status of tasks.

		You can also check other projects at https://github.com/PecheslavVenik/backend-projects`,
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewStatusDoneCmd())
	cmd.AddCommand(NewStatusInProgressCmd())
	cmd.AddCommand(NewStatusTodoCmd())

	return cmd
}
