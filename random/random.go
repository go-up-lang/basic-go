package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Account struct
type RandomNumber struct {
	Rand_Num  int
	Bonus_Num int
}

func Random() []RandomNumber {
	rand.Seed(time.Now().UnixNano())
	var randomNumbers []RandomNumber

	for i := 0; i < 5; i++ {
		randNum := rand.Intn(41) + 5
		bonusNum := rand.Intn(41) + 5
		randomNumbers = append(randomNumbers, RandomNumber{
			Rand_Num:  randNum,
			Bonus_Num: bonusNum,
		})
	}
	fmt.Println("이 값은 random.go : ", randomNumbers)
	return randomNumbers
}
