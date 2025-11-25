package xbase62

import (
	"crypto/rand"
	"math/big"
)

// Rand generates a random base62 string of the specified length using alphanumeric characters
func Rand(length int) string {
	return randString(length, Alphanumeric)
}

// RandNumeric generates a random numeric string of the specified length
func RandNumeric(length int) string {
	return randString(length, Numeric)
}

// RandAlphabetic generates a random alphabetic string of the specified length
func RandAlphabetic(length int) string {
	return randString(length, Alphabetic)
}

func randString(length int, charset string) string {
	if length <= 0 {
		return ""
	}

	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range length {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			panic(err)
		}
		result[i] = charset[num.Int64()]
	}

	return string(result)
}
