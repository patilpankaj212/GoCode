package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
The program accepts two command line arguments
1. name of hotel
2. total bill amount

calculate and print the tip amount for the bill amount.
tip is 10% of total bill
*/

func main() {

	// os.Args is a string slice that contains all the arguments
	// first argument would be the application binary which we would skip
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Incorrect Arguments passed!!! \nProgram accepts two aruguments \n 1. Name of Hotel \n 2. Total bill amount")
	} else {
		nameOfHotel := args[0]
		billAmout, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Name of Hotel: ", nameOfHotel)
		fmt.Println("Total Bill amount: ", billAmout)
		fmt.Printf("tip amount: %.2f\n", calculateTip(billAmout))
	}
}

func calculateTip(billAmout float64) float64 {
	if billAmout <= 0 {
		return 0
	}
	return 0.1 * billAmout
}
