package main

import ("fmt"

"github.com/serin0837/learngo/accounts"

"log"
)


func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(),account.ChangeOwner("serin"))
}
