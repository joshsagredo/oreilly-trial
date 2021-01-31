package random

import (
	"crypto/rand"
	"math/big"
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
	//r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	res := make([]byte, randomLength)
	for i := 0; i < randomLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			return "", err
		}
		res[i] = all[num.Int64()]
	}

	return string(res), nil
}
