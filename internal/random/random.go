package random

import (
	"math/rand"
	"strings"
)

var (
	chars    = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	digits   = "0123456789"
	specials = "=+*/!@#$?"
	all      = string(chars) + digits + specials
)

// GenerateUsername generates a random string of numbers and characters
func GenerateUsername(length int) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// GeneratePassword generates a random ASCII string with at least one digit and one special character.
func GeneratePassword(length int) string {
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

// PickEmail picks a random item from emailDomains slice
func PickEmail(emailDomains []string) string {
	return emailDomains[rand.Intn(len(emailDomains))]
}
