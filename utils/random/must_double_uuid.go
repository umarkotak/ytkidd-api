package random

import (
	"strings"

	"github.com/google/uuid"
)

// it will generate and join N uuid
func MustGenUUIDTimes(times int) string {
	uuids := []string{}

	if times <= 0 {
		times = 1
	}

	for i := 1; i <= times; i++ {
		uuid, _ := uuid.NewRandom()
		uuids = append(uuids, uuid.String())
	}

	return strings.Join(uuids, "-")
}
