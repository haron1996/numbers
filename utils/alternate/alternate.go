package alternate

import (
	"math/rand"
	"time"

	"github.com/haron1996/numbers/utils/setOne"
	//"github.com/haron1996/numbers/utils/setThree"
	//"github.com/haron1996/numbers/utils/setTwo"
)

func GeneratePhoneNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch r.Intn(3) {
	case 0:
		return setOne.GenerateStartingWith0700710720740And79()
	case 1:
		//return setTwo.GenerateStartingWith0757075807590768And0769()
		return setOne.GenerateStartingWith0700710720740And79()
	default:
		//return setThree.GenerateStartingWith01100111011201130114And0115()
		return setOne.GenerateStartingWith0700710720740And79()
	}
}
