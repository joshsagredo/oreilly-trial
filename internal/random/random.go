package random

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
)

const (
	Chars        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
	Digits       = "0123456789"
	Specials     = "=+*/!@#$?"
	All          = Chars + Digits + Specials
	TypeUsername = "TYPE_USERNAME"
	TypePassword = "TYPE_PASSWORD"
)

/*// GenerateUsername generates a random string of numbers and characters
func GenerateUsername(length int) (string, error) {
	/*rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()


}*/

// Generate generates a random string for username or password
func Generate(length int, outputType string) (string, error) {
	var sourceChars string

	switch outputType {
	case TypeUsername:
		sourceChars = Chars
	case TypePassword:
		sourceChars = All
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

/*// GeneratePassword generates a random ASCII string with at least one digit and one special character.
func GeneratePassword(length int) (string, error) {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(All))))
		if err != nil {
			return "", err
		}

		ret[i] = All[num.Int64()]
	}

	return string(ret), nil
}*/

// PickEmail picks a random item from emailDomains slice
func PickEmail(emailDomains []string) string {
	return emailDomains[mathRand.Intn(len(emailDomains))]
}
