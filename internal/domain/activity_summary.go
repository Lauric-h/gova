package domain

type ActivitySummary struct {
	TotalDistance float32
	TotalAscent   int
	TotalDuration int
	SportType     SportType
	Period        Period
	Count         int
}

func (a *ActivitySummary) GetDistanceInKm() float32 {
	return a.TotalDistance / 1000
}

func (a *ActivitySummary) GetDurationInHours() float32 {
	return float32(a.TotalDuration / 3600)
}
