package random

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGenerateUsername(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cases := []struct{
		caseName string
		randomLength int
	}{
		{"random10", 10},
		{"random20", 20},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			var b strings.Builder
			for i := 0; i < tc.randomLength; i++ {
				b.WriteRune(chars[rand.Intn(len(chars))])
			}

			t.Logf("username generated. username=%s\n", b.String())
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cases := []struct{
		caseName string
		randomLength int
	}{
		{"random10", 10},
		{"random20", 20},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			buf := make([]byte, tc.randomLength)
			buf[0] = digits[rand.Intn(len(digits))]
			buf[1] = specials[rand.Intn(len(specials))]
			for i := 2; i < tc.randomLength; i++ {
				buf[i] = all[rand.Intn(len(all))]
			}

			rand.Shuffle(len(buf), func(i, j int) {
				buf[i], buf[j] = buf[j], buf[i]
			})

			t.Logf("password generated. password=%s\n", string(buf))
		})
	}
}