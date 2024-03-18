package setOne

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateStartingWith0700710720740And79() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefixes := []string{"070", "071", "072", "074", "079"}
	randomPrefix := prefixes[r.Intn(len(prefixes))]
	randomSuffix := 100000000 + r.Intn(900000000)
	randomSuffixStr := fmt.Sprintf("%09d", randomSuffix)
	phoneNumber := fmt.Sprintf("%s%s", randomPrefix, randomSuffixStr)
	phoneNumber = phoneNumber[:10]
	return phoneNumber
}
