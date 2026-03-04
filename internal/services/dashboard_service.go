package services

type TodayView struct {
	Habits []HabitView
}

func GetToday() (*TodayView, error) {

	habits, err := ListHabits()
	if err != nil {
		return nil, err
	}

	return &TodayView{
		Habits: habits,
	}, nil
}
