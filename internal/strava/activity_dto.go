package strava

type Activity struct {
	Distance  float64 `json:"distance"`
	Duration  int     `json:"elapsed_time"`
	Ascent    float64 `json:"total_elevation_gain"`
	SportType string  `json:"sport_type"`
}
