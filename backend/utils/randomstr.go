package utils

import "math/rand"

var letters = "abcdefghjkmnopqrstuvwxyz1234567890"

func RandomPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
