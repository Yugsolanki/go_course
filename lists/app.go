package main

import "fmt"

func main() {
	prices := [4]float64{1.3, 2.5, 3.7, 4.3}
	fmt.Println(prices)

	var productName [4]string
	fmt.Println(productName)
	// productName = [4]string{"bla", "blaa"}
	productName[2] = "second book"
	fmt.Println(productName)

	// Slices
	featuredPrices := prices[1:3] // index 1 is included and index 3 is excluded. [2.5, 3.7]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	featuredPrices[0] = 100.99 // this will also change the `prices` array as it is not a copy but reference
	fmt.Println(featuredPrices)
	fmt.Println("OG Prices Array: ", prices)

	highlightedPrices := featuredPrices[:1] // this is a slice of a slice
	fmt.Println(highlightedPrices)

	fmt.Println("Length: ", len(prices), " Capacity: ", cap(prices))
	fmt.Println("Length: ", len(featuredPrices), " Capacity: ", cap(featuredPrices))
	fmt.Println("Length: ", len(highlightedPrices), " Capacity: ", cap(highlightedPrices))
	// length represents the number of elements in an slice or array
	// capacity in `highlightedPrices` is 3 because `featuredPrices` is a slice of `prices` and it has three elements to it
	// in `highlightedPrices` we have only selected 1 but there are two to the right which can be selected too by [:2], [:3]
	// therefore the capacity is 3, but in `featuredPrices` we have already selected all the element to the right since we went [1:]
	// so there are no elements left to the right so therefore it's capacity is 3 only

	featuredPrices = prices[0:3]
	// look here, we have one more element to right of the `prices` array which could be selected but we excluded it
	// so the capacity of `featuredPrices` will be 4
	fmt.Println("Length: ", len(featuredPrices), "Capacity: ", cap(featuredPrices))
	// so you can always select more items to the right (end of the array) but not to the left

}

// ------------ DYNAMIC ARRAY ---------------

// func main2() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	updatedPrices := append(prices, 5.99)
// 	// prices := append(prices, 5.99)
// 	fmt.Println(prices, updatedPrices)

// 	// removing a element
// 	prices = prices[1:]

// 	// Unpacking values
// 	discountPrices := []float64{100.33, 42.34, 23.23, 73.32}
// 	prices = append(prices, discountPrices...)
// 	fmt.Println(prices)
// }
