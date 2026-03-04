package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vikas-bhat-d/flow-cli/internal/services"
)

var habitCmd = &cobra.Command{
	Use:   "habit",
	Short: "Manage Your habits",
}

var habitAddCmd = &cobra.Command{
	Use:  "add [name]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		services.AddHabit(args[0])

		fmt.Println("Habit added:", args[0])
	},
}
var habitListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

		habits, err := services.ListHabits()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, h := range habits {

			status := "[ ]"
			if h.DoneToday {
				status = "[x]"
			}

			fmt.Printf("%s [%d] %s 🔥 %d\n",
				status,
				h.ID,
				h.Name,
				h.CurrentStreak,
			)
		}
	},
}

var habitDoneCmd = &cobra.Command{
	Use:  "done [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := strconv.Atoi(args[0])

		services.CompleteHabit(id)

		fmt.Println("Habit completed")
	},
}

func init() {
	rootCmd.AddCommand(habitCmd)
	habitCmd.AddCommand(habitAddCmd)
	habitCmd.AddCommand(habitListCmd)
	habitCmd.AddCommand(habitDoneCmd)
}
