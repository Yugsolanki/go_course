package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	arr := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&arr, double)
	tripled := transformNumbers(&arr, triple)
	fmt.Println(doubled, tripled)

	// Function as Return Values
	// transformerFn := getTransformerFunction("double") can also pass this down
	doubled = transformNumbers(&arr, getTransformerFunction("double"))

	// Anonymous Functions
	quadraple := transformNumbers(&arr, func(number int) int {
		return number * 4
	})
	fmt.Println(quadraple)

	// Closures
	seven := transformNumbers(&arr, createTransformer(7))
	fmt.Println(seven)

	eight := transformNumbers(&arr, createTransformer(8))
	fmt.Println(eight)

	// Recursion
	fac := factorial(5)
	fmt.Println(fac)

	// Variadic Function
	sum := sumUp(0, 1, 2, 3, 4, 5)
	fmt.Println(sum)
	// when we have a defined variable like `arr` and want to pass to a variadic function,
	// you can do this `arr...` to open it up
	sum = sumUp(1, arr...)
	fmt.Println(sum)
}

func transformNumbers(arr *[]int, transform transformFn) []int {
	doubled := []int{}

	for _, val := range *arr {
		doubled = append(doubled, transform(val))
	}

	return doubled
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// function as return type
func getTransformerFunction(str string) transformFn {
	if str == "double" {
		return double
	} else {
		return triple
	}
}

// Closure:
// A closure is a function that captures the environment in which it was created.
// In this case, the closure captures the variable 'factor' from the parent function 'createTransformer'.
// This allows the returned function to access 'factor' even after 'createTransformer' has finished
// we can use factor here even though it is not part of the anonymous func, but since it is part of the parent function we can
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Recursion
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Variadic Function
func sumUp(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
