package domain

import "errors"

type SportType string

const (
	SportTypeTrailRun       SportType = "TrailRun"
	SportTypeRun            SportType = "run"
	SportTypeHike           SportType = "Hike"
	SportTypeRide           SportType = "Ride"
	SportTypeWeightTraining SportType = "WeightTraining"
)

func (s SportType) String() string {
	return string(s)
}

func SportTypeFromString(i string) (SportType, error) {
	switch SportType(i) {
	case SportTypeTrailRun, SportTypeRun, SportTypeHike, SportTypeRide, SportTypeWeightTraining:
		return SportType(i), nil
	default:
		return SportType(""), errors.New("invalid SportType")
	}
}
