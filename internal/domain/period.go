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

func CreateWeek(shouldGetLast bool) Period {
	now := time.Now()

	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	start := now.AddDate(0, 0, -(weekday - 1))
	if shouldGetLast {
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

func CreateMonth(shouldGetLast bool) Period {
	now := time.Now()

	if shouldGetLast {
		now = now.AddDate(0, -1, 0)
	}

	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	nextMonth := now.AddDate(0, 1, 0)
	nextMonthFirstDay := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())
	end := nextMonthFirstDay.Add(-time.Second)

	return Period{
		StartDay: start,
		EndDay:   end,
	}
}
