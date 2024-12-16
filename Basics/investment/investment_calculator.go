package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, expectedReturnRate, years float64
	const inflation = 2.5
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years) // getting the input from the user from the terminal
	var futureValue float64 = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	var inflatedValue float64 = futureValue / math.Pow(1+inflation/100, years)

	str1 := fmt.Sprintf("Future value: %v\n", futureValue)
	str2 := fmt.Sprintf("Future inflated value: %v\n", inflatedValue)
	fmt.Print(str1, str2)

}
