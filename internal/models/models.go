package models

type HabitLog struct {
	HabitID int    `json:"habit_id"`
	Date    string `json:"date"`
}

type Habit struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CurrentStreak int    `json:"current_streak"`
	LongestStreak int    `json:"longest_streak"`
}

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	HabitID int    `json:"habit_id,omitempty"`

	Done     bool   `json:"done"`
	DoneDate string `json:"done_date,omitempty"`

	CreatedAt    string `json:"created_at"`
	ScheduledFor string `json:"scheduled_for,omitempty"`

	Deadline string `json:"deadline,omitempty"`

	Estimate int `json:"estimate,omitempty"`
	Spent    int `json:"spent,omitempty"`
}

type Session struct {
	Duration int    `json:"duration"`
	TaskID   int    `json:"task_id,omitempty"`
	Date     string `json:"date"`
}
