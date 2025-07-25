package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func worker[T int, K int](inputChan <-chan T, outputChan chan<- K) {
// 	go func() {
// 		for val := range inputChan {
// 			time.Sleep(2 * time.Second)
// 			outputChan <- K(val * 2)
// 		}
// 	}()
// }

// func main() {
// 	inputChan := make(chan int)
// 	outputChan := make(chan int)

// 	randomGenerator := func() int { return rand.Intn(500) }

// 	// Start 10 workers
// 	for i := 0; i < 3; i++ {
// 		go worker(inputChan, outputChan)
// 	}

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			inputChan <- randomGenerator()
// 		}
// 		close(inputChan)
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			result := <-outputChan
// 			fmt.Println(result)
// 		}
// 		close(outputChan)
// 	}()

// 	time.Sleep(100 * time.Second)
// }
