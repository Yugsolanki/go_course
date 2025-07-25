package main

// import (
// 	"fmt"
// 	"sync"
// )

// var (
// 	counter int
// 	mutex   sync.Mutex
// )

// func SafeIncrement(wg *sync.WaitGroup, id int) {
// 	defer wg.Done()

// 	for i := 0; i < 1000; i++ {
// 		mutex.Lock()
// 		counter++
// 		mutex.Unlock()
// 	}

// 	fmt.Printf("Goroutine %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go SafeIncrement(&wg, i)
// 	}

// 	wg.Wait()
// 	fmt.Println("Final counter value:", counter)
// }
