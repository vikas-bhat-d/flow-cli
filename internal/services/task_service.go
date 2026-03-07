package services

import (
	"time"

	"github.com/vikas-bhat-d/flow-cli/internal/models"
	"github.com/vikas-bhat-d/flow-cli/internal/storage"
)

func AddTask(
	title string,
	habitID int,
	scheduledFor string,
	deadline string,
	estimate int,
) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	now := time.Now()

	created := now.Format("2006-01-02")

	if scheduledFor == "" {
		scheduledFor = created
	}

	task := models.Task{
		ID:           len(state.Tasks) + 1,
		Title:        title,
		HabitID:      habitID,
		Done:         false,
		CreatedAt:    created,
		ScheduledFor: scheduledFor,
		Deadline:     deadline,
		Estimate:     estimate,
		Spent:        0,
	}

	state.Tasks = append(state.Tasks, task)

	return storage.SaveState(state)
}

type TaskView struct {
	ID           int
	Title        string
	Done         bool
	HabitName    string
	ScheduledFor string
	Deadline     string
	Estimate     int
	Spent        int
}

func ListTasks(today bool, date string, pending bool, done bool) ([]TaskView, error) {

	state, err := storage.LoadState()
	if err != nil {
		return nil, err
	}

	todayDate := time.Now().Format("2006-01-02")

	var tasks []TaskView

	for _, t := range state.Tasks {

		// Filter: today
		if today && t.ScheduledFor != todayDate {
			continue
		}

		// Filter: date
		if date != "" && t.ScheduledFor != date {
			continue
		}

		// Filter: pending
		if pending && t.Done {
			continue
		}

		// Filter: done
		if done && !t.Done {
			continue
		}

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
			ID:           t.ID,
			Title:        t.Title,
			Done:         t.Done,
			HabitName:    habitName,
			ScheduledFor: t.ScheduledFor,
			Estimate:     t.Estimate,
			Spent:        t.Spent,
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

func AddFocusTime(taskID int, minutes int) error {

	state, err := storage.LoadState()
	if err != nil {
		return err
	}

	for i := range state.Tasks {

		if state.Tasks[i].ID == taskID {

			state.Tasks[i].Spent += minutes

			if state.Tasks[i].Estimate > 0 &&
				state.Tasks[i].Spent >= state.Tasks[i].Estimate {

				state.Tasks[i].Done = true
				state.Tasks[i].DoneDate = time.Now().Format("2006-01-02")

				if state.Tasks[i].HabitID != 0 {
					updateHabitProgress(state, state.Tasks[i].HabitID)
				}
			}

			break
		}
	}

	return storage.SaveState(state)
}
