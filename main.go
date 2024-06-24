package main

import (
	"fmt"

	"github.com/my-project/random"
)

func main() {
	lotto := random.Random()

	fmt.Println("로또 번호는 : ", lotto)
}
