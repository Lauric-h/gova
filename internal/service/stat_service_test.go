package service

import (
	"gova/internal/core"
	"gova/internal/domain"
	"testing"
	"time"
)

type MockApiClient struct {
	Activities []core.Activity
	Err        error
}

func (m *MockApiClient) ListActivities(before int64, after int64) ([]core.Activity, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Activities, nil
}

func (m *MockApiClient) GetCurrentAthlete() (*core.Athlete, error) {
	return nil, nil // pas utilisé dans ce test
}

func TestListActivities_AggregatesBySportType(t *testing.T) {
	mockClient := &MockApiClient{
		Activities: []core.Activity{
			{Distance: 5000, Duration: 1800, Ascent: 100, SportType: "Run"},
			{Distance: 10000, Duration: 3600, Ascent: 200, SportType: "Run"},
			{Distance: 8000, Duration: 2700, Ascent: 500, SportType: "TrailRun"},
		},
	}

	service := NewStatService(mockClient)
	period := domain.Period{
		StartDay: time.Now().AddDate(0, 0, -7),
		EndDay:   time.Now(),
	}

	result, err := service.ListActivities(period)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 sport types, got %d", len(result))
	}

	run, ok := result["Run"]
	if !ok {
		t.Fatal("Expected Run in results")
	}
	if run.Count != 2 {
		t.Errorf("Expected 2 runs, got %d", run.Count)
	}
	if run.TotalDistance != 15000 {
		t.Errorf("Expected total distance 15000, got %f", run.TotalDistance)
	}
	if run.TotalDuration != 5400 {
		t.Errorf("Expected total duration 5400, got %d", run.TotalDuration)
	}
	if run.TotalAscent != 300 {
		t.Errorf("Expected total ascent 300, got %d", run.TotalAscent)
	}

	trail, ok := result["TrailRun"]
	if !ok {
		t.Fatal("Expected TrailRun in results")
	}
	if trail.Count != 1 {
		t.Errorf("Expected 1 trail run, got %d", trail.Count)
	}
}

func TestListActivities_IgnoresInvalidSportTypes(t *testing.T) {
	mockClient := &MockApiClient{
		Activities: []core.Activity{
			{Distance: 5000, Duration: 1800, Ascent: 100, SportType: "Run"},
			{Distance: 3000, Duration: 900, Ascent: 0, SportType: "Yoga"},
			{Distance: 2000, Duration: 600, Ascent: 0, SportType: "Swimming"},
		},
	}

	service := NewStatService(mockClient)
	period := domain.Period{
		StartDay: time.Now().AddDate(0, 0, -7),
		EndDay:   time.Now(),
	}

	result, err := service.ListActivities(period)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 sport type, got %d", len(result))
	}

	if _, ok := result["Run"]; !ok {
		t.Error("Expected Run in results")
	}
}
