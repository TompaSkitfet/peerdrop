package util

import "math/rand"

func GenerateId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
