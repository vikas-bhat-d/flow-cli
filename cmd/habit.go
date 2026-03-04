package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vikas-bhat-d/flow-cli/internal/storage"
)

var habitCmd = &cobra.Command{
	Use:   "habit",
	Short: "Manage Your habits",
}

var habitAddCmd = &cobra.Command{
	Use:  "add [name]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		state, _ := storage.LoadState()

		newHabit := storage.Habit{
			ID:   len(state.Habits) + 1,
			Name: args[0],
		}

		state.Habits = append(state.Habits, newHabit)

		storage.SaveState(state)

		fmt.Println("Habit added:", args[0])
	},
}

var habitListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

		state, _ := storage.LoadState()

		for _, h := range state.Habits {
			fmt.Printf("[%d] %s 🔥 %d\n",
				h.ID,
				h.Name,
				h.CurrentStreak,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(habitCmd)
	habitCmd.AddCommand(habitAddCmd)
	habitCmd.AddCommand(habitListCmd)
}