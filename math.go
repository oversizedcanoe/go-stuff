package main

import (
	"fmt"
	"strconv"
)

func doMath() {
	firstNumber := promptNumber("first")
	secondNumber := promptNumber("second")

	sum := firstNumber + secondNumber

	fmt.Println("Sum = " + strconv.Itoa(sum))
}

func promptNumber(ordinalPositionName string) int {
	var input string
	var result int

	for {
		fmt.Println("Enter " + ordinalPositionName + " number:")
		fmt.Scanln(&input)
		number, err := strconv.Atoi(input)

		// Can also do in one line:
		// if number, err := strconv.Atoi(input); err != nil {
		if err != nil {
			fmt.Println("Bad input '" + input + "'. Try again")
		} else {
			result = number
			break
		}

	}
	return result
}
