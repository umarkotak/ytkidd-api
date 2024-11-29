package utils

import "time"

func CalculateAge(birthTime, currentTime time.Time) int64 {
	yearsDifference := int64(currentTime.Sub(birthTime).Hours() / 24 / 365.25)
	return yearsDifference
}
