package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 // default value is assigned 0.0
	var expectedReturnRate float64
	var years float64

	// var investmentAmount = 1000
	// var investmentAmount float64 = 1000
	// expectedReturnRate := 5.5 let go infer type itself
	// var investAmount, years float64 = 1000, 10
	// var investAmount, years = 1000, 10
	// var investAmount, years float64 := 1000, 10
	// var investAmount, years := 1000.0, 10.0

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	// var futureValue = (float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))) with type conversion
	// var futureValue = (investmentAmount * math.Pow(1+expectedReturnRate/100, years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %v \nFuture Value (adjusted for inflation): %v", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.0f \nFuture Value (adjusted for inflation): %.0f \n", futureValue, futureRealValue) // rounded value, no decimal
	fmt.Printf("Future Value: %.1f \nFuture Value (adjusted for inflation): %.1f \n", futureValue, futureRealValue) // one decimal value

	// formattedString := fmt.Sprintf("Future Value: %.1f \nFuture Value (adjusted for inflation): %.1f \n", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for inflation)", futureRealValue)
}

// func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
// 	futureValue := (investmentAmount * math.Pow(1+expectedReturnRate/100, years))
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
// 	return futureValue, futureRealValue
// }

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = (investmentAmount * math.Pow(1+expectedReturnRate/100, years))
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	// return futureValue, futureRealValue
	return
}
