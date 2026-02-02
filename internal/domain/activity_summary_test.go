package domain

import "testing"

func TestActivitySummary_GetDistanceInKm(t *testing.T) {
	activitySummary := ActivitySummary{
		TotalDistance: 100000,
		TotalAscent:   6000,
		TotalDuration: 36000,
		SportType:     SportTypeTrailRun,
		Period:        CreateWeek(false),
		Count:         1,
	}

	if activitySummary.GetDistanceInKm() != 100 {
		t.Errorf("Expected 100 got %f", activitySummary.GetDistanceInKm())
	}
}

func TestActivitySummary_GetDurationInHours(t *testing.T) {
	activitySummary := ActivitySummary{
		TotalDistance: 100000,
		TotalAscent:   6000,
		TotalDuration: 36000,
		SportType:     SportTypeTrailRun,
		Period:        CreateWeek(false),
		Count:         1,
	}

	if activitySummary.GetDurationInHours() != 10 {
		t.Errorf("Expected 10 got %f", activitySummary.GetDurationInHours())
	}
}
