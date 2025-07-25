package main

import "fmt"

// func main() {

// 	myChannel := make(chan string)
// 	anotherChannel := make(chan string)

// 	go func() {
// 		fmt.Println("got it")
// 		myChannel <- "Hello"
// 	}()

// 	go func() {
// 		fmt.Println("got it 2")
// 		anotherChannel <- "World"
// 	}()

// 	select {
// 	case msg := <-myChannel:
// 		fmt.Println(msg)
// 	case msg := <-anotherChannel:
// 		fmt.Println(msg)
// 	}
// }

// func main() {
// 	myChannel := make(chan string, 3)
// 	chars := []string{"a", "b", "c"}

// 	for _, s := range chars {
// 		myChannel <- s
// 	}

// 	close(myChannel)

// 	for result := range myChannel {
// 		fmt.Println(result)
// 	}
// }

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
