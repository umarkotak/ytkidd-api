package utils

import "time"

func RemainingDays(future time.Time) int64 {
	now := time.Now()
	// Zero out time parts to compare only dates
	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	futureDate := time.Date(future.Year(), future.Month(), future.Day(), 0, 0, 0, 0, future.Location())

	days := futureDate.Sub(nowDate).Hours() / 24
	return int64(days)
}
