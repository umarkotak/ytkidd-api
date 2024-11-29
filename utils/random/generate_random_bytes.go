package random

import (
	"crypto/rand"
)

func GenBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
