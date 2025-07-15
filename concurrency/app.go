package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan) tell go when to stop looking for chan
}

func main() {
	// simply using `go greet("nice to meet you")` won't actually print anything
	// as the job of main function is the simply dispatch the go function
	// meaning it simply triggers the execution of the function but dosn't wait for the result
	// it move on after trigger the execution, so `main` starts all func and end itself and dosn't wait for their output

	// channel
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How .... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course", done)

	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// }

	for range done {

	}
}
