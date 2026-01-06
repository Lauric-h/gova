package domain

import "time"

type Period struct {
	StartDay time.Time
	EndDay   time.Time
}

const (
	PeriodWeek  = "week"
	PeriodMonth = "month"
)

func CreateWeek(ShouldGetLast bool) Period {
	now := time.Now()

	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	start := now.AddDate(0, 0, -(weekday - 1))
	if ShouldGetLast {
		start = start.AddDate(0, 0, -7)
	}

	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())

	end := start.AddDate(0, 0, 6)
	end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), end.Location())

	return Period{
		StartDay: start,
		EndDay:   end,
	}
}

func CreateMonth(ShouldGetLast bool) Period {
	return Period{}
}
