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

func Random() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		Random_value := rand.Intn(45) + 1
		Bonus_value := rand.Intn(45) + 1
		clova := RandomNumber{Rand_Num: Random_value, Bonus_Num: Bonus_value}
		fmt.Println(clova)
	}

}
