package accounts

import ("errors"
	"fmt"
)

type Account struct{
	owner string
	balance int
}

//NewAccount create Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance:0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int){
	a.balance += amount
}
//Balance of your account
func (a Account) Balance() int{
	return a.balance
}
//Withdraw x amount 
func(a *Account) Withdraw(amount int) error{
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
	//error value
}
//ChangeOwner
func (a *Account) ChangeOwner(newOwner string){
	a.owner = newOwner
}
//Owner
func (a Account) Owner() string{
	return a.owner
}

func(a Account) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\nHas:", a.Balance())
}