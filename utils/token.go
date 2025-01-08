package utils

import (
	"math/rand"
	"time"
)

const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

/*
* GenerateToken generates a random string of the specified length
* if length is passed in as 0, then the token is of length 16
 */
func GenerateToken(length int) string {
	if length == 0 {
		length = 16
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // Create a new random generator
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = CHARSET[r.Intn(len(CHARSET))]
	}
	return string(result)
}
