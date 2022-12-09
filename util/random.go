package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	// golang的字串拼接: https://www.cnblogs.com/apocelipes/p/9283841.html

	var sb strings.Builder
	len := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generate a random email string
func RandomEmail() string {
	var mailSb strings.Builder

	mailSb.WriteString(RandomString(5))
	mailSb.WriteString("@example.com")

	return mailSb.String()
}
