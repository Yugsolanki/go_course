package main

// import (
// 	"fmt"
// 	"sync"
// )

// var once sync.Once

// var message string

// func initialize() {
// 	fmt.Println("Initializing...")
// 	message = "Hello World"
// }

// func getMessage() string {
// 	// once.Do(passed function will only run once)
// 	// and then once will be set to true
// 	// so the function will not run again
// 	once.Do(initialize)
// 	return message
// }

// func main() {
// 	fmt.Println(getMessage())
// 	fmt.Println(getMessage())
// 	fmt.Println(getMessage())
// }
