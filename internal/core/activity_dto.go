package core

type Activity struct {
	Distance  float64 `json:"distance"`
	Duration  int     `json:"elapsed_time"`
	Ascent    float64 `json:"total_elevation_gain"`
	SportType string  `json:"sport_type"`
}

type Athlete struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Sex       string `json:"sex"`
	Premium   bool   `json:"premium"`
}
