package random

import (
	"math/rand"
	"time"
)

// Account struct
type RandomNumber struct {
	Rand_Num  int
	Bonus_Num int
}

func Random() RandomNumber {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(41) + 5
	bonusNum := rand.Intn(41) + 5

	return RandomNumber{
		Rand_Num:  randNum,
		Bonus_Num: bonusNum,
	}

}
