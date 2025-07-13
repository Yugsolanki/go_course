package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)

	result2 := add(3.4, 5.5)
	fmt.Println(result2)

	result3 := add("blaa", " huuhhh")
	fmt.Println(result3)
}

// func add[T any](a, b T) T {
// 	return a + b
// }

func add[T int | float64 | string](a, b T) T {
	return a + b
}
