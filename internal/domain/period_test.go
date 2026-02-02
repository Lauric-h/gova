package domain

import (
	"testing"
	"time"
)

func TestCreateWeek_CurrentWeek(t *testing.T) {
	period := CreateWeek(false)

	// Start should be a Monday
	if period.StartDay.Weekday() != time.Monday {
		t.Errorf("Expected start day to be Monday, got %s", period.StartDay.Weekday())
	}

	// End should be a Sunday
	if period.EndDay.Weekday() != time.Sunday {
		t.Errorf("Expected end day to be Sunday, got %s", period.EndDay.Weekday())
	}

	// Start should be at 00:00:00
	if period.StartDay.Hour() != 0 || period.StartDay.Minute() != 0 || period.StartDay.Second() != 0 {
		t.Errorf("Expected start time to be 00:00:00, got %02d:%02d:%02d",
			period.StartDay.Hour(), period.StartDay.Minute(), period.StartDay.Second())
	}

	// End should be at 23:59:59
	if period.EndDay.Hour() != 23 || period.EndDay.Minute() != 59 || period.EndDay.Second() != 59 {
		t.Errorf("Expected end time to be 23:59:59, got %02d:%02d:%02d",
			period.EndDay.Hour(), period.EndDay.Minute(), period.EndDay.Second())
	}

	// Period should span approximately 7 days (Mon 00:00:00 to Sun 23:59:59)
	duration := period.EndDay.Sub(period.StartDay)
	minDuration := 7*24*time.Hour - 2*time.Second
	maxDuration := 7 * 24 * time.Hour
	if duration < minDuration || duration > maxDuration {
		t.Errorf("Expected duration of ~7 days, got %v", duration)
	}

	// Current time should be within the period
	now := time.Now()
	if now.Before(period.StartDay) || now.After(period.EndDay) {
		t.Errorf("Current time %v should be within period [%v, %v]",
			now, period.StartDay, period.EndDay)
	}
}

func TestCreateWeek_LastWeek(t *testing.T) {
	currentWeek := CreateWeek(false)
	lastWeek := CreateWeek(true)

	// Last week should be exactly 7 days before current week
	expectedStart := currentWeek.StartDay.AddDate(0, 0, -7)
	if !lastWeek.StartDay.Equal(expectedStart) {
		t.Errorf("Expected last week start to be %v, got %v", expectedStart, lastWeek.StartDay)
	}

	expectedEnd := currentWeek.EndDay.AddDate(0, 0, -7)
	if !lastWeek.EndDay.Equal(expectedEnd) {
		t.Errorf("Expected last week end to be %v, got %v", expectedEnd, lastWeek.EndDay)
	}

	// Last week should also be Monday to Sunday
	if lastWeek.StartDay.Weekday() != time.Monday {
		t.Errorf("Expected last week start to be Monday, got %s", lastWeek.StartDay.Weekday())
	}
	if lastWeek.EndDay.Weekday() != time.Sunday {
		t.Errorf("Expected last week end to be Sunday, got %s", lastWeek.EndDay.Weekday())
	}
}

func TestCreateMonth_CurrentMonth(t *testing.T) {
	period := CreateMonth(false)
	now := time.Now()

	// Start should be the 1st of current month
	if period.StartDay.Day() != 1 {
		t.Errorf("Expected start day to be 1st, got %d", period.StartDay.Day())
	}
	if period.StartDay.Month() != now.Month() {
		t.Errorf("Expected start month to be %s, got %s", now.Month(), period.StartDay.Month())
	}
	if period.StartDay.Year() != now.Year() {
		t.Errorf("Expected start year to be %d, got %d", now.Year(), period.StartDay.Year())
	}

	// Start should be at 00:00:00
	if period.StartDay.Hour() != 0 || period.StartDay.Minute() != 0 || period.StartDay.Second() != 0 {
		t.Errorf("Expected start time to be 00:00:00, got %02d:%02d:%02d",
			period.StartDay.Hour(), period.StartDay.Minute(), period.StartDay.Second())
	}

	// End should be the last day of current month at 23:59:59
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	expectedLastDay := nextMonth.Add(-time.Second).Day()
	if period.EndDay.Day() != expectedLastDay {
		t.Errorf("Expected end day to be %d, got %d", expectedLastDay, period.EndDay.Day())
	}

	if period.EndDay.Hour() != 23 || period.EndDay.Minute() != 59 || period.EndDay.Second() != 59 {
		t.Errorf("Expected end time to be 23:59:59, got %02d:%02d:%02d",
			period.EndDay.Hour(), period.EndDay.Minute(), period.EndDay.Second())
	}

	// Current time should be within the period
	if now.Before(period.StartDay) || now.After(period.EndDay) {
		t.Errorf("Current time %v should be within period [%v, %v]",
			now, period.StartDay, period.EndDay)
	}
}

func TestCreateMonth_LastMonth(t *testing.T) {
	currentMonth := CreateMonth(false)
	lastMonth := CreateMonth(true)

	// Last month should be the month before
	expectedMonth := currentMonth.StartDay.Month() - 1
	expectedYear := currentMonth.StartDay.Year()
	if expectedMonth == 0 {
		expectedMonth = 12
		expectedYear--
	}

	if lastMonth.StartDay.Month() != expectedMonth {
		t.Errorf("Expected last month to be %s, got %s", expectedMonth, lastMonth.StartDay.Month())
	}
	if lastMonth.StartDay.Year() != expectedYear {
		t.Errorf("Expected last month year to be %d, got %d", expectedYear, lastMonth.StartDay.Year())
	}

	// Start should be the 1st
	if lastMonth.StartDay.Day() != 1 {
		t.Errorf("Expected start day to be 1st, got %d", lastMonth.StartDay.Day())
	}

	// End should be the last day of last month
	if lastMonth.EndDay.Month() != expectedMonth {
		t.Errorf("Expected end month to be %s, got %s", expectedMonth, lastMonth.EndDay.Month())
	}
}

func TestCreateMonth_HandlesVariableMonthLengths(t *testing.T) {
	// This test verifies the period calculation works correctly
	// by checking that end day is always valid for the month
	period := CreateMonth(false)

	// End day should be between 28 and 31
	endDay := period.EndDay.Day()
	if endDay < 28 || endDay > 31 {
		t.Errorf("Expected end day to be between 28-31, got %d", endDay)
	}

	// The day after EndDay should be the 1st of next month
	nextDay := period.EndDay.Add(time.Second)
	if nextDay.Day() != 1 {
		t.Errorf("Expected day after end to be 1st, got %d", nextDay.Day())
	}
}
