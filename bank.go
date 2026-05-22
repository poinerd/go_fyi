package main

import(
	"fmt"
	"sync"
)


type account struct{
	mu sync.RWMutex
	balance int 
	id int
}

var accounts []*account

//  same := []int{}

func getAllAccounts (){
	for i:= 0 ; i < len(accounts) ; i++{
		fmt.Println(accounts[i].balance)
	}
}

func createAccount () *account{
    
	newAccount := &account{
		balance : 0,
		id : len(accounts) + 1,
	}
	accounts = append(accounts, newAccount)
	return newAccount
}


func (a *account) deposit(amount int, wg *sync.WaitGroup) int{
    defer wg.Done()
	a.mu.Lock()
    
	init := a.balance
	a.balance = init + amount

    defer a.mu.Unlock()
	return a.balance
}

func (a *account) getBalance() {
	a.mu.RLock()
    defer a.mu.RUnlock()

    fmt.Printf("Your account has the balance of %d ", a.balance)
	
}

func main(){
	newAcc := createAccount()

	var wg sync.WaitGroup
	wg.Add(4)

	go newAcc.deposit(500, &wg)
	go newAcc.deposit(500, &wg)
	go newAcc.deposit(500, &wg)
	go newAcc.deposit(500, &wg)

	wg.Wait()
	newAcc.getBalance()

	// getAllAccounts()


}