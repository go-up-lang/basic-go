package main

import "fmt"

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5)
}
