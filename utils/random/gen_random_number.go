package random

import (
	"math"
	"math/rand"
)

// it will generate N digit number
func GenNDigitNumber(length int) int64 {
	if length <= 0 {
		length = 1
	}

	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length)) - 1

	randomNumber := rand.Int63n(max-min+1) + min

	return randomNumber
}
