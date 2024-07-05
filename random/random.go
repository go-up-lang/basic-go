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
		for j := 1; j <= 5; j++ {
			Random_value := rand.Intn(45) + 1
			fmt.Print(Random_value, ", ")
		}
		Bonus_value := rand.Intn(45) + 1
		fmt.Println(" + ", Bonus_value)
	}

}
