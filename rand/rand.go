package rand

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// CreateRandomBytes generate rand []byte by specify chars.
func CreateRandomBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// RandInt return random int between [min max]
func RandInt(min, max int) int {
	if max == min {
		return min
	}
	if max < min {
		max, min = min, max
	}

	r.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond * 1)
	return min + r.Intn(max-min)
}
