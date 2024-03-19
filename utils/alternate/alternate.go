package alternate

import (
	"math/rand"
	"time"

	"github.com/haron1996/numbers/utils/setOne"
	"github.com/haron1996/numbers/utils/setThree"
	"github.com/haron1996/numbers/utils/setTwo"
	"github.com/haron1996/numbers/utils/zeroSevenTen"
)

func GeneratePhoneNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch r.Intn(4) {
	case 0:
		return setOne.GenerateStartingWith0700710720740And79()
	case 1:
		return setTwo.GenerateStartingWith0757075807590768And0769()
	case 2:
		return setThree.GenerateStartingWith01100111011201130114And0115()
	default:
		return zeroSevenTen.Generate0710()
	}
}
