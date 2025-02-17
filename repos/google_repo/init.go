package google_repo

import (
	"crypto/rsa"
)

var googlePublicKeys map[string]*rsa.PublicKey

func Initialize() {
	// Initialize the public key cache.  In a real application, you should
	// refresh these keys periodically.
	googlePublicKeys = make(map[string]*rsa.PublicKey)
}
