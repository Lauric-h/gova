package service

import (
	"fmt"
	"gova/internal/client"
)

type StatService struct {
	Client *client.Client
}

func NewStatService(client *client.Client) *StatService {
	return &StatService{Client: client}
}

func (s *StatService) ListActivities(shouldGetLast bool) {
	fmt.Printf("last %t", shouldGetLast)
	activities, err := s.Client.ListActivities(1767552000, 1766956800)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(activities)
}
