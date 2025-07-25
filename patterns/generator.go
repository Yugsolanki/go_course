package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
// 	stream := make(chan T)

// 	go func() {
// 		defer close(stream)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case stream <- fn():
// 			}
// 		}
// 	}()

// 	return stream
// }

// func main() {
// 	done := make(chan int)

// 	defer close(done)

// 	randomGenerator := func() int { return rand.Intn(50000000) }

// 	for rando := range repeatFunc(done, randomGenerator) {
// 		fmt.Println(rando)
// 		if rando < 100 {
// 			done <- 1
// 		}
// 	}
// }
