package main

import (
	fileops "concurrency/Basics/bank/fileops"
	"fmt"
	"os"
)

const fileName = "balance.txt"

func main() {

	accountBalance, err := fileops.ReadFromFile(fileName)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Welcome to the Bank of Golang!")

	for {
		fmt.Println("Enter 1 to check your account balance")
		fmt.Println("Enter 2 to deposit money")
		fmt.Println("Enter 3 to withdraw money")
		fmt.Println("Enter 4 to Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current account balance is: %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter your deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Printf("Deposit successful. New balance: %.2f\n", accountBalance)
				fileops.WriteBalanceToFile(accountBalance, fileName)
			} else {
				fmt.Println("Sorry, Invalid your money deposit.")
			}
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter your withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount > 0 && withdrawalAmount <= accountBalance {
				accountBalance -= withdrawalAmount
				fmt.Printf("Withdrawal successful. New balance: %.2f\n", accountBalance)
				fileops.WriteBalanceToFile(accountBalance, fileName)
			} else {
				fmt.Println("Sorry, Invalid your money withdrawal.")
			}
		default:
			fmt.Println("Thanks for choosing Bank of Golang!")
			return
		}
	}

}
