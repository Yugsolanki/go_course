package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// func worker(inputChan <-chan string, outputChan chan<- string) {
// 	go func() {
// 		for word := range inputChan {
// 			time.Sleep(2 * time.Second)
// 			outputChan <- strings.ToUpper(word)
// 		}
// 	}()
// }

// func main() {
// 	inputChan := make(chan string)
// 	outputChan := make(chan string)

// 	// Initialize 3 workers
// 	for i := 0; i < 3; i++ {
// 		go worker(inputChan, outputChan)
// 	}

// 	// Sending data to input channel
// 	go func() {
// 		words := [12]string{"apple", "banana", "cherry", "date", "berry", "orange", "dragon", "mango", "straw", "watermelon", "litchi", "melon"}
// 		for _, word := range words {
// 			inputChan <- word
// 		}
// 		close(inputChan)
// 	}()

// 	// Reading data out of output channel
// 	go func() {
// 		for i := 0; i < 12; i++ {
// 			result := <-outputChan
// 			fmt.Println(result)
// 		}
// 	}()

// 	time.Sleep(10 * time.Second)
// }
