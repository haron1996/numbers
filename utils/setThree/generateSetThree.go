package setThree

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateStartingWith01100111011201130114And0115() string {
	prefixes := []string{"0110", "0111", "0112", "0113", "0114", "0115"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := prefixes[r.Intn(len(prefixes))]
	randomNumber := rand.Intn(1000000)
	randomNumberString := fmt.Sprintf("%06d", randomNumber)
	phoneNumber := fmt.Sprintf("%s%s", prefix, randomNumberString)
	phoneNumber = phoneNumber[:10]
	return phoneNumber
}
