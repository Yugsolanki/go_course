package main

// func main() {
// 	done := make(chan bool)
// 	counter := 0
// 	times := 2

// 	for i := 0; i < times; i++ {
// 		go func(i int) {
// 			fmt.Println("Goroutine initiated:", i)
// 			counter++
// 			fmt.Println("Sending data to channel")
// 			done <- true
// 			fmt.Println("Data sent successfully")
// 		}(i)
// 	}

// 	for i := 0; i < times; i++ {
// 		fmt.Println("\nAbout to receive data")
// 		<-done
// 		fmt.Println("Received data succesfully")
// 	}
// }

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		message := <-ch
// 		fmt.Println("Received:", message)
// 	}()

// 	fmt.Println("About to send")
// 	ch <- "Hello"
// 	fmt.Println("Sent")
// }

// func main() {
// 	ch := make(chan int, 2)

// 	ch <- 1
// 	ch <- 2
// 	fmt.Println("Sending third item now")
// 	ch <- 3 // blocks since the buffer is full
// 	fmt.Println("this line never executes")
// }

// func main() {
// 	ch := make(chan int, 2)

// 	ch <- 1
// 	ch <- 2

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println("About to receive for the third time but from an empty channel")
// 	fmt.Println(<-ch)
// }

// func main() {
// 	ch := make(chan int, 2)

// 	// Filling the buffer
// 	ch <- 1
// 	ch <- 2

// 	// try sending without blocking
// 	select {
// 	case ch <- 3:
// 		fmt.Println("Successfully sent 3")
// 	default:
// 		fmt.Println("Channel is full, cannot send anything more")
// 	}

// 	// Emptying channel to hit the default below
// 	// <-ch
// 	// <-ch

// 	// Try to receive without blocking
// 	select {
// 	case value := <-ch:
// 		fmt.Println("Received: ", value)
// 	default:
// 		fmt.Println("Channel is empty, nothing to receive.")
// 	}
// }

// func badProducer(ch chan<- string) {
// 	ch <- "Hello"
// 	// close(ch) // if not closed it will throw error in receiver's side
// }

// func main() {
// 	ch := make(chan string)

// 	go badProducer(ch)

// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}

// 	fmt.Println("If not closed it will throw error before reach this print")
// }
