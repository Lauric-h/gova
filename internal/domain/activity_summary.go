package domain

type ActivitySummary struct {
	TotalDistance int
	TotalAscent   int
	TotalDuration int
	SportType     SportType
	Period        Period
	Count         int
}
