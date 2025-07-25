package main

// import (
// 	"fmt"
// 	"sync"
// )

// type BankAccount struct {
// 	balance int
// 	mutex   sync.Mutex
// }

// func (b *BankAccount) deposit(wg *sync.WaitGroup, amount int) {
// 	defer wg.Done()

// 	b.mutex.Lock()
// 	defer b.mutex.Unlock()

// 	fmt.Println("Depositing:", amount)
// 	b.balance += amount
// 	fmt.Println("New balance after deposit:", b.balance)
// }

// func (b *BankAccount) withdraw(wg *sync.WaitGroup, amount int) {
// 	defer wg.Done()

// 	b.mutex.Lock()
// 	defer b.mutex.Unlock()

// 	fmt.Println("Withdrawing:", amount)
// 	b.balance -= amount
// 	fmt.Println("New balance after withdrawal:", b.balance)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	var account = BankAccount{
// 		balance: 500,
// 	}

// 	wg.Add(3)
// 	go account.deposit(&wg, 500)
// 	go account.withdraw(&wg, 500)
// 	go account.deposit(&wg, 500)

// 	// you can also avoid passing `wg` to function by doing this
// 	// wg.Add(1)
// 	// go func ()  {
// 	// 	defer wg.Done()
// 	// 	account.deposit(500)
// 	// }()

// 	wg.Wait()
// 	fmt.Println("Final balance:", account.balance)
// }
