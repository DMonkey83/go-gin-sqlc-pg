// Package util ...
package util

import (
	"math/rand"
	"strings"
	"time"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

// init ...
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt Generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString Generate a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(abc)

	for i := 0; i < n; i++ {
		c := abc[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner ...
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney ...
func RandomMoney() int64 {
	return RandomInt(0, 994)
}

// RandomCurrency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP", "JPY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
