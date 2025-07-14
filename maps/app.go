package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"]) // yes it is case-sensitive bro

	// Adding items
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	// Deleting
	delete(websites, "Google")
	fmt.Println(websites)

	fmt.Println("-------------------------------------------------------------")

	// Make
	userNames := make([]string, 2, 5)
	// makes an array of length 2, and those 2 elements are empty
	// 5 is the capacity, it tells go the allocate enough space for 5 elements
	// but create an array with two empty slots
	// when working with `[]string{}` the issue is you cannot do this `userNames[0] = 2`
	// but `make` will create an array with specified empty slot (2 in this case) and then we can access it like `userNames[0]`
	fmt.Println(userNames) // [ ]
	userNames = append(userNames, "Yug")
	userNames = append(userNames, "Solanki")
	// these two will be appended to the end, so now the length will be 4
	fmt.Println(userNames)
	userNames[0] = "first"
	userNames[1] = "second"
	fmt.Println(userNames)

	fmt.Println("-------------------------------------------------------------")

	// Map Make
	courseRatings := make(map[string]float64, 2) // cannot define capacity
	fmt.Println(courseRatings)
	courseRatings["go"] = 4.6
	courseRatings["react"] = 3.6
	courseRatings["js"] = 4.7 // will have to re-alocate memory for this
	fmt.Println(courseRatings)

	fmt.Println("-------------------------------------------------------------")

	// Type alias
	appRatings := make(floatMap, 2)
	appRatings["chrome"] = 4.6
	appRatings["insta"] = 3.6
	appRatings.output()

	fmt.Println("-------------------------------------------------------------")

	// For loop
	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:",key,"Value:",value)
	}
}
