package main

import "fmt"

func main() {
	age := 32

	// var agePointer *int
	// agePointer := &age
	// fmt.Println("Age address: ", agePointer)
	// fmt.Println("Value at address: ", *agePointer)

	adultYears := getAdultYears(&age)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	// *age = *age - 18 this will mutate the age variable
	return *age - 18
}
