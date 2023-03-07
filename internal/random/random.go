package random

import (
	mathRand "math/rand"
	"time"
)

const (
	chars        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
	digits       = "0123456789"
	specials     = "=+*/!@#$?"
	all          = chars + digits + specials
	randomLength = 12
)

// GeneratePassword generates a random string for username or password
func GeneratePassword() (string, error) {
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	res := make([]byte, randomLength)
	for i := 0; i < randomLength; i++ {
		num := r.Intn(len(all))

		res[i] = all[num]
	}

	return string(res), nil
}
