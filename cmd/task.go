package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vikas-bhat-d/flow-cli/internal/services"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var habitID int
var scheduledFor string
var deadline string
var estimate int

var today bool
var date string
var pending bool
var done bool

var taskAddCmd = &cobra.Command{
	Use:  "add [title]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := services.AddTask(
			args[0],
			habitID,
			scheduledFor,
			deadline,
			estimate,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("✓ Task added:", args[0])
	},
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := services.ListTasks(today, date, pending, done)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, t := range tasks {

			status := "[ ]"
			if t.Done {
				status = "[x]"
			}

			progress := ""
			if t.Estimate > 0 {
				progress = fmt.Sprintf(" (%d/%dm)", t.Spent, t.Estimate)
			}

			if t.HabitName != "" {

				fmt.Printf(
					"%s [%d] %s%s (habit: %s)\n",
					status,
					t.ID,
					t.Title,
					progress,
					t.HabitName,
				)

			} else {

				fmt.Printf(
					"%s [%d] %s%s\n",
					status,
					t.ID,
					t.Title,
					progress,
				)
			}
		}
	},
}

var taskDoneCmd = &cobra.Command{
	Use:  "done [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = services.CompleteTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("✓ Task completed")
	},
}

func init() {

	rootCmd.AddCommand(taskCmd)

	taskAddCmd.Flags().IntVar(&habitID, "habit", 0, "Link task to a habit")

	taskAddCmd.Flags().StringVar(
		&scheduledFor,
		"date",
		"",
		"Schedule task for date (YYYY-MM-DD)",
	)

	taskAddCmd.Flags().StringVar(
		&deadline,
		"deadline",
		"",
		"Deadline date (YYYY-MM-DD)",
	)

	taskAddCmd.Flags().IntVar(
		&estimate,
		"estimate",
		0,
		"Estimated effort in minutes",
	)

	taskListCmd.Flags().BoolVar(&today, "today", false, "Show tasks scheduled today")

	taskListCmd.Flags().StringVar(
		&date,
		"date",
		"",
		"Show tasks scheduled for date (YYYY-MM-DD)",
	)

	taskListCmd.Flags().BoolVar(
		&pending,
		"pending",
		false,
		"Show only pending tasks",
	)

	taskListCmd.Flags().BoolVar(
		&done,
		"done",
		false,
		"Show completed tasks",
	)

	taskCmd.AddCommand(taskAddCmd)
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskDoneCmd)
}
