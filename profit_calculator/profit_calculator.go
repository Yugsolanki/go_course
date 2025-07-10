package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Enter Tax Rate: ")
	// fmt.Scan(&taxRate)

	revenue, err := readInput("Revenue")
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return
	}

	expenses, err := readInput("Expenses")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := readInput("Tax Rate")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)
	storeResults(earningsBeforeTax, earningsAfterTax, ratio)
}

func readInput(text string) (float64, error) {
	fmt.Printf("Enter %v: ", text)
	var temp float64
	fmt.Scan(&temp)

	if temp <= 0 {
		return 0, errors.New("Invalid input, the value should be greater than 0.")
	}

	return temp, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
