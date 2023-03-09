package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabitc = "abcdefghigklmnopqrsntuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInteger(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabitc)
	for i := 0; i < n; i++ {
		c := alphabitc[rand.Intn(k)]

		sb.WriteByte(c)
	}
	return sb.String()
}

func RnadomName() string {
	return RandomString(7)
}

func RandomPrice() int {
	return int(RandomInteger(50, 500))
}

func RandomEmail() string {
	return RandomString(7) + "@gmail.com"
}
