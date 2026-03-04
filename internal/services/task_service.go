package services

import (
	"time"

	"github.com/vikas-bhat-d/flow-cli/internal/models"
	"github.com/vikas-bhat-d/flow-cli/internal/storage"
)

func AddTask(title string, habitID int) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	task := models.Task{
		ID:      len(state.Tasks) + 1,
		Title:   title,
		HabitID: habitID,
		Done:    false,
	}

	state.Tasks = append(state.Tasks, task)

	return storage.SaveState(state)
}

type TaskView struct {
	ID        int
	Title     string
	Done      bool
	HabitName string
}

func ListTasks() ([]TaskView, error) {

	state, err := storage.LoadState()
	if err != nil {
		return nil, err
	}

	var tasks []TaskView

	for _, t := range state.Tasks {

		habitName := ""

		if t.HabitID != 0 {
			for _, h := range state.Habits {
				if h.ID == t.HabitID {
					habitName = h.Name
					break
				}
			}
		}

		tasks = append(tasks, TaskView{
			ID:        t.ID,
			Title:     t.Title,
			Done:      t.Done,
			HabitName: habitName,
		})
	}

	return tasks, nil
}

func updateHabitProgress(state *storage.State, habitID int) {

	today := time.Now().Format("2006-01-02")

	count := 0

	for _, t := range state.Tasks {

		if t.HabitID == habitID &&
			t.Done &&
			t.DoneDate == today {

			count++
		}
	}

	for i := range state.Habits {

		if state.Habits[i].ID == habitID {

			// if at least one task today → habit is active
			if count == 1 {

				state.Habits[i].CurrentStreak++

				if state.Habits[i].CurrentStreak > state.Habits[i].LongestStreak {
					state.Habits[i].LongestStreak = state.Habits[i].CurrentStreak
				}
			}

			break
		}
	}
}

func CompleteTask(id int) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	today := time.Now().Format("2006-01-02")

	for i := range state.Tasks {

		if state.Tasks[i].ID == id {

			if state.Tasks[i].Done {
				return nil
			}

			state.Tasks[i].Done = true
			state.Tasks[i].DoneDate = today

			habitID := state.Tasks[i].HabitID

			if habitID != 0 {
				updateHabitProgress(state, habitID)
			}

			break
		}
	}

	return storage.SaveState(state)
}
