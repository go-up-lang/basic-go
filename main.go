package main

import (
	"fmt"

	"github.com/my-project/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.owner = "new owner"
	fmt.Println(account)
}
