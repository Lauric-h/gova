package domain

import "time"

type ActivitySummary struct {
	TotalDistance int
	TotalAscent   int
	TotalDuration int
	SportType     SportType
	StartDay      time.Time
	EndDay        time.Time
	Count         int
}
