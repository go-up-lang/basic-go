package main

import "fmt"

func ifmoon(age int) bool {

	if koreanAge := age + 2; koreanAge >= 18 {
		return true
	}

	return false

}
func main() {
	fmt.Println(ifmoon(16))
}
