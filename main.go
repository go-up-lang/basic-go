package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		randNum := rand.Intn(96) + 5
		fmt.Println(randNum)
	}
}
