package random

import (
	"crypto/rand"
	"errors"
	"math/big"
	mathRand "math/rand"
	"time"
)

const (
	Chars        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
	Digits       = "0123456789"
	Specials     = "=+*/!@#$?"
	All          = Chars + Digits + Specials
	TypeUsername = "TYPE_USERNAME"
	TypePassword = "TYPE_PASSWORD"
)

func init() {
	mathRand.Seed(time.Now().Unix())
}

// Generate generates a random string for username or password
func Generate(length int, outputType string) (string, error) {
	var sourceChars string

	switch outputType {
	case TypeUsername:
		sourceChars = Chars
	case TypePassword:
		sourceChars = All
	}

	if length > 32 {
		return "", errors.New("invalid random value")
	}

	res := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(sourceChars))))
		if err != nil {
			return "", err
		}

		res[i] = sourceChars[num.Int64()]
	}

	return string(res), nil
}

// PickEmail picks a random item from emailDomains slice
func PickEmail(emailDomains []string) string {
	return emailDomains[mathRand.Intn(len(emailDomains))]
}
