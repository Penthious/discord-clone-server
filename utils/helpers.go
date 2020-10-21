package utils

import (
	"fmt"
	"math/rand"
)

func GetRandomString(prefix string, length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(charset[rand.Intn(len(charset))])
	}

	return fmt.Sprintf("%s_%s", prefix, string(runes))
}
