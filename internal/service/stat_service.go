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

func (s *StatService) ListActivities(period domain.Period) (map[string]domain.ActivitySummary, error) {
	//activities, err := s.Client.ListActivities(1767552000, 1766956800)
	activities, err := s.Client.ListActivities(period.EndDay.Unix(), period.StartDay.Unix())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	formattedActivities := make(map[string]domain.ActivitySummary)
	for _, act := range activities {
		sportType, err := domain.SportTypeFromString(act.SportType)
		if err != nil {
			// TODO log
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
