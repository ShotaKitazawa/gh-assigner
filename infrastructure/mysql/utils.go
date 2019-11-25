package mysql

import "time"

type Timestamp struct {
	RequestedAt *time.Time `db:"requested_at"`
	ReviewedAt  *time.Time `db:"reviewed_at"`
}

func calcDuration(timestamps []Timestamp) (duration time.Duration) {
	for _, timestamp := range timestamps {
		if timestamp.RequestedAt == nil || timestamp.ReviewedAt == nil {
			continue
		}
		duration += timestamp.ReviewedAt.Sub(*timestamp.RequestedAt)
	}
	return
}
