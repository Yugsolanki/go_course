package main

// import "math/rand"

// func generator[K any, T int](done <-chan K, fn func() T) <-chan T {
// 	generatorStream := make(chan T)

// 	go func() {
// 		defer close(generatorStream)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case generatorStream <- fn():
// 			}
// 		}
// 	}()

// 	return generatorStream
// }

// func square[T int, K any](done <-chan K, generator <-chan T) <-chan T {
// 	squaresStream := make(chan T)

// 	go func() {
// 		defer close(squaresStream)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case squaresStream <- (<-generator * 2):
// 			}
// 		}
// 	}()

// 	return squaresStream
// }

// func even[T int, K any](done <-chan K, squareStream <-chan T) <-chan bool {
// 	even := make(chan bool)

// 	go func() {
// 		defer close(even)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case even <- (<-squareStream%3 == 0):
// 			}
// 		}
// 	}()

// 	return even
// }

// func main() {
// 	done := make(chan bool)

// 	randomGenerator := func() int { return rand.Intn(500) }

// 	for isEven := range even(done, square(done, generator(done, randomGenerator))) {
// 		println(isEven)
// 		// if !isEven {
// 		// 	done <- false
// 		// }
// 	}
// }
