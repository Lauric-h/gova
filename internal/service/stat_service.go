package service

import (
	"fmt"
	"gova/internal/core"
	"gova/internal/domain"
)

type StatService struct {
	Client core.ApiClient
}

func NewStatService(client core.ApiClient) *StatService {
	return &StatService{Client: client}
}

func (s *StatService) GetAthleteSummary() (*domain.AthleteSummary, error) {
	athlete, err := s.Client.GetCurrentAthlete()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch athlete summary: %w", err)
	}

	return &domain.AthleteSummary{
		Username:  athlete.Username,
		Firstname: athlete.Firstname,
		Lastname:  athlete.Lastname,
		City:      athlete.City,
		State:     athlete.State,
		Country:   athlete.Country,
		Sex:       athlete.Sex,
		Premium:   athlete.Premium,
	}, nil
}

func (s *StatService) ListActivities(period domain.Period) (map[string]domain.ActivitySummary, error) {
	activities, err := s.Client.ListActivities(period.EndDay.Unix(), period.StartDay.Unix())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	formattedActivities := make(map[string]domain.ActivitySummary)
	for _, act := range activities {
		sportType, err := domain.SportTypeFromString(act.SportType)
		if err != nil {
			continue
		}

		if entry, ok := formattedActivities[sportType.String()]; ok {
			entry.TotalDistance += float32(act.Distance)
			entry.TotalAscent += int(act.Ascent)
			entry.TotalDuration += act.Duration
			entry.Count++

			formattedActivities[sportType.String()] = entry

			continue
		}

		formattedActivities[sportType.String()] = domain.ActivitySummary{
			TotalDistance: float32(act.Distance),
			TotalAscent:   int(act.Ascent),
			TotalDuration: act.Duration,
			SportType:     sportType,
			Count:         1,
		}
	}

	return formattedActivities, nil
}
