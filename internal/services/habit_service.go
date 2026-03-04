package services

import (
	"time"

	"github.com/vikas-bhat-d/flow-cli/internal/models"
	"github.com/vikas-bhat-d/flow-cli/internal/storage"
)

type HabitView struct {
	ID            int
	Name          string
	CurrentStreak int
	DoneToday     bool
}

func AddHabit(name string) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	habit := models.Habit{
		ID:   len(state.Habits) + 1,
		Name: name,
	}

	state.Habits = append(state.Habits, habit)

	return storage.SaveState(state)
}

func CompleteHabit(id int) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	today := time.Now().Format("2006-01-02")

	for _, log := range state.HabitLogs {
		if log.HabitID == id && log.Date == today {
			return nil
		}
	}

	state.HabitLogs = append(state.HabitLogs, models.HabitLog{
		HabitID: id,
		Date:    today,
	})

	for i, h := range state.Habits {

		if h.ID == id {

			state.Habits[i].CurrentStreak++

			if state.Habits[i].CurrentStreak > state.Habits[i].LongestStreak {
				state.Habits[i].LongestStreak = state.Habits[i].CurrentStreak
			}
		}
	}

	return storage.SaveState(state)
}

func ListHabits() ([]HabitView, error) {

	state, err := storage.LoadState()
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")

	var habits []HabitView

	for _, h := range state.Habits {

		doneToday := false

		for _, log := range state.HabitLogs {
			if log.HabitID == h.ID && log.Date == today {
				doneToday = true
				break
			}
		}

		habits = append(habits, HabitView{
			ID:            h.ID,
			Name:          h.Name,
			CurrentStreak: h.CurrentStreak,
			DoneToday:     doneToday,
		})
	}

	return habits, nil
}
