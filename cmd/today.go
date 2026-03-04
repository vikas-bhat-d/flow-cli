package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vikas-bhat-d/flow-cli/internal/services"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's productivity dashboard",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := services.GetToday()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Today")
		fmt.Println()

		fmt.Println("Habits")

		if len(data.Habits) == 0 {
			fmt.Println("No habits yet.")
			return
		}

		for _, h := range data.Habits {

			status := "[ ]"
			if h.DoneToday {
				status = "[x]"
			}

			fmt.Printf("%s %s 🔥 %d\n",
				status,
				h.Name,
				h.CurrentStreak,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
