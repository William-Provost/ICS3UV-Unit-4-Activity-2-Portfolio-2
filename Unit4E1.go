// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-27
// Fileoverview: This program prints the Fibonacci sequence up to a user-entered limit.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// variable declaration
	var limitString string
	var limit int
	var firstNumber int = 0
	var secondNumber int = 1
	var nextNumber int

	reader := bufio.NewReader(os.Stdin)

	// get limit from user
	fmt.Print("Enter the limit for Fibonacci series: ")
	limitString, _ = reader.ReadString('\n')
	limitString = strings.TrimSpace(limitString)
	limit, _ = strconv.Atoi(limitString)

	fmt.Println("")
	fmt.Println(firstNumber)
	fmt.Println(secondNumber)

	// while loop to generate Fibonacci numbers
	nextNumber = firstNumber + secondNumber

	for nextNumber <= limit {
		fmt.Println(nextNumber)

		firstNumber = secondNumber
		secondNumber = nextNumber
		nextNumber = firstNumber + secondNumber
	}

	fmt.Println("\nDone.")
}
