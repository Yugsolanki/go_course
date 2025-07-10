package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func main() {
	var accountBalance, err = readBalanceFromFile()

	if err != nil {
		fmt.Print("ERROR: ")
		fmt.Println(err)
		fmt.Println("-------------------------")
		// return to stop running of the code
		// panic("Can't continue, sorry.") even this will exit the code
	}

	fmt.Println("Welcome to Go Bank!")

	// switch choice {
	// case 1:
	// 	fmt.Println("Choice 1")
	// case 2:
	// 	fmt.Println("Choice 2")
	// default:
	// 	fmt.Println("Default")
	// 	return
	// }

	for {
		fmt.Println("\n-------------------------")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your account balance is: %.2f", accountBalance)
		} else if choice == 2 {
			var amountToDeposit float64
			fmt.Printf("Enter the amount to deposit: ")
			fmt.Scan(&amountToDeposit)

			if amountToDeposit <= 0 {
				fmt.Println("Invalid amount. Amount must be greater than 0.")
				continue
			}

			accountBalance += amountToDeposit
			fmt.Printf("Balance updated: %.2f", accountBalance)
		} else if choice == 3 {
			var amountToWithdraw float64
			fmt.Printf("Enter the amount to withdraw: ")
			fmt.Scan(&amountToWithdraw)

			if amountToWithdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if amountToWithdraw > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw most than you have.")
				continue
			}

			accountBalance -= amountToWithdraw
			fmt.Printf("Balance updated: %.2f", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			writeBalanceToFile(accountBalance)
			break
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}
