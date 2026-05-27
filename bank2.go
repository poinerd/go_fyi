package main 

import(
  "fmt"
  "sync"
)

// This study exposes me to goroutines, waitgroups and mutexes
//I still have a very serious misunderstanding of pointers and how they are being used here.

type account struct{
	mu sync.RWMutex
	id int
	balance int
}

var accounts []*account // thi stores coordinats of the accounts created

func getAllAccount () []*account{
	return accounts
}

func createAccount () *account{

	newAccount := &account{
		id : len(accounts) + 1,
		balance : 0,
	}

	accounts = append(accounts, newAccount)
	return newAccount
}

func (a *account) deposit (amount int, wg *sync.WaitGroup) int{
	defer wg.Done()

	a.mu.Lock()
	defer a.mu.Unlock()

	initialBalance := a.balance
	a.balance = initialBalance + amount
	return a.balance

}

func (a *account) getBalance()int{
	a.mu.RLock()
	defer a.mu.RUnlock()

	accountBalance := a.balance
	return accountBalance
}

func main(){
	newUser := createAccount()
    var wg sync.WaitGroup
	wg.Add(4)

	go newUser.deposit(500, &wg)
	go newUser.deposit(500, &wg)
	go newUser.deposit(500, &wg)
	go newUser.deposit(500, &wg)

	wg.Wait()

	balance := newUser.getBalance()-
	fmt.Println(balance)
}


//waitgroup.Add() waitgroup.Wait() waitgroup.Done()