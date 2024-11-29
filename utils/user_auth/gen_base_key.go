package user_auth

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mathRand "math/rand"
)

func GenBaseKey() {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)

	publicKeyString := hex.EncodeToString(publicKey)
	privateKeyString := hex.EncodeToString(privateKey)
	josePrivateKey := generate32DigitNumber()

	fmt.Printf("jwt private key: %v\n", privateKeyString)
	fmt.Printf("jwt public key: %v\n", publicKeyString)
	fmt.Printf("jwt secret key: %v\n", josePrivateKey)
}

func generate32DigitNumber() string {
	var digits [32]byte
	for i := range digits {
		digits[i] = byte(mathRand.Intn(10) + '0') // Generate a random digit (0-9) and convert to ASCII
	}
	return string(digits[:]) // Convert the byte slice to a string
}
