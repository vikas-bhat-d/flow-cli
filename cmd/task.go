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

var taskAddCmd = &cobra.Command{
	Use:  "add [title]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := services.AddTask(args[0], habitID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("✓ Task added:", args[0])
	},
}

var taskListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := services.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, t := range tasks {

			status := "[ ]"
			if t.Done {
				status = "[x]"
			}

			if t.HabitName != "" {
				fmt.Printf("%s [%d] %s (habit: %s)\n",
					status,
					t.ID,
					t.Title,
					t.HabitName,
				)
			} else {
				fmt.Printf("%s [%d] %s\n",
					status,
					t.ID,
					t.Title,
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
	taskCmd.AddCommand(taskAddCmd)
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskDoneCmd)
}
