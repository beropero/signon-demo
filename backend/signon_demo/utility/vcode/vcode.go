package vcode

import (
	"math/rand"
	"time"
)

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}