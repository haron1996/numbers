package zeroSevenTen

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate0710() string {
	prefixes := []string{"0710"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := prefixes[r.Intn(len(prefixes))]
	remainingDigits := ""
	for i := 0; i < 6; i++ {
		remainingDigits += fmt.Sprintf("%d", rand.Intn(10))
	}
	phoneNumber := fmt.Sprintf("%s%s", prefix, remainingDigits)
	phoneNumber = phoneNumber[:10]
	return phoneNumber
}
