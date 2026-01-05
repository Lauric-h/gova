package service

import (
	"fmt"
	"gova/internal/client"
	"gova/internal/domain"
)

type StatService struct {
	Client *client.Client
}

func NewStatService(client *client.Client) *StatService {
	return &StatService{Client: client}
}

func (s *StatService) ListActivities(shouldGetLast bool) (map[string]domain.ActivitySummary, error) {
	activities, err := s.Client.ListActivities(1767552000, 1766956800)
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
			entry.TotalDistance += int(act.Distance)
			entry.TotalAscent += int(act.Ascent)
			entry.TotalDuration += act.Duration
			entry.Count++

			continue
		}

		formattedActivities[sportType.String()] = domain.ActivitySummary{
			TotalDistance: int(act.Distance),
			TotalAscent:   int(act.Ascent),
			TotalDuration: int(act.Duration),
			SportType:     sportType,
			Count:         1,
		}
	}

	return formattedActivities, nil
}
