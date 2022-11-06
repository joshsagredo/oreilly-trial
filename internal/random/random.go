package random

import (
	"crypto/rand"
	"errors"
	"math/big"
	mathRand "math/rand"
	"time"
)

const (
	Chars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
	Digits   = "0123456789"
	Specials = "=+*/!@#$?"
	All      = Chars + Digits + Specials
)

func init() {
	mathRand.Seed(time.Now().Unix())
}

// GeneratePassword generates a random string for username or password
func GeneratePassword(length int) (string, error) {
	if length < 8 || length > 32 {
		return "", errors.New("invalid random value")
	}

	res := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(All)))) //nolint:gosec
		if err != nil {
			return "", err
		}

		res[i] = All[num.Int64()]
	}

	return string(res), nil
}
