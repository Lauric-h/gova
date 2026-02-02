package domain

import "testing"

func TestSportType_String(t *testing.T) {
	sport := SportTypeTrailRun
	if sport.String() != "TrailRun" {
		t.Errorf("Expected TrailRun got %s", sport.String())
	}
}

func TestSportType_FromString(t *testing.T) {
	s := "TrailRun"
	sport, _ := SportTypeFromString(s)
	if sport != SportTypeTrailRun {
		t.Errorf("Expected TrailRun got %s", sport.String())
	}
}
