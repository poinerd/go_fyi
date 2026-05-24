package main

import(
	"fmt"
)

type account struct{
	id int
	balance int
}

var accounts []*account

// var accounta := []int {}
// *account this means a pointer to the account type


func createAccount() *account{
	newAccount := &account{
		id : len(accounts) + 1,
		balance : 0,
	}
	
	accounts = append(accounts, newAccount)
	return newAccount

}

func (a *account) depositAccount(d int) int{
	
	fmt.Println(*a)
	initialBalance := a.balance
	a.balance = initialBalance + d
	return a.balance
}

func main(){
	account1 := createAccount()
	fmt.Println(account1.depositAccount(56))

}