package task

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

type Status string

const (
	ToDo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"title"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      ToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status Status) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case ToDo:
		for _, t := range tasks {
			if t.Status == ToDo {
				filteredTasks = append(filteredTasks, t)
			}
		}
	case InProgress:
		for _, t := range tasks {
			if t.Status == InProgress {
				filteredTasks = append(filteredTasks, t)
			}
		}
	case Done:
		for _, t := range tasks {
			if t.Status == Done {
				filteredTasks = append(filteredTasks, t)
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Description", "Created At", "Updated At"})

	table.SetAutoWrapText(true)
	table.SetColWidth(50)

	for _, task := range filteredTasks {
		statusIndicator := "●"
		statusColor := tablewriter.FgWhiteColor
		switch task.Status {
		case ToDo:
			statusColor = tablewriter.FgHiBlackColor
		case InProgress:
			statusColor = tablewriter.FgRedColor
		case Done:
			statusColor = tablewriter.FgGreenColor
		}
		table.Rich([]string{
			fmt.Sprintf("%d", task.ID),
			statusIndicator,
			task.Description,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"),
		}, []tablewriter.Colors{
			{}, {tablewriter.Bold, statusColor}, {}, {}, {},
		})
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	table.Render()

	return nil
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66"))

	formattedId := style.Render(fmt.Sprintf("(ID: %d)", task.ID))
	fmt.Printf("\nTask added successfully: %s\n\n", formattedId)
	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(tasks) {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask deleted successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskStatus(id int64, status Status) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case ToDo:
				task.Status = ToDo
			case InProgress:
				task.Status = InProgress
			case Done:
				task.Status = Done
			}
			task.UpdatedAt = time.Now()
		}

		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}
