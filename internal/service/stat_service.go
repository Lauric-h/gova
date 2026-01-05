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

func (s *StatService) ListActivities(shouldGetLast bool) (*domain.ActivitySummary, error) {
	activities, err := s.Client.ListActivities(1767552000, 1766956800)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var activity domain.ActivitySummary
	for _, act := range activities {
		fmt.Println(act.SportType)
		sportType, err := domain.SportTypeFromString(act.SportType)
		fmt.Println(sportType)
		if err != nil {
			// TODO log
			continue
		}
		activity.SportType = sportType

		activity.TotalDistance += int(act.Distance)
		activity.TotalAscent += int(act.Ascent)
		activity.TotalDuration += act.Duration
	}

	return &activity, nil
}
