package main

import (
	"fmt"

	"example.com/bank/utils"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {
	var accountBalance, err = utils.ReadFloatFromFile(fileName)

	if err != nil {
		fmt.Print("ERROR: ")
		fmt.Println(err)
		fmt.Println("-------------------------")
		// return to stop running of the code
		// panic("Can't continue, sorry.") even this will exit the code
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())

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
		presentOptions()

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
			utils.WriteFloatToFile(accountBalance, fileName)
			break
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}
