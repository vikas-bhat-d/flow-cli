package storage

import "github.com/vikas-bhat-d/flow-cli/internal/models"

type State struct {
	Version   int
	Habits    []models.Habit
	Tasks     []models.Task
	Sessions  []models.Session
	HabitLogs []models.HabitLog
}
