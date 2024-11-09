package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	
	temp := number
	numDigits := 0
	for temp > 0 {
		numDigits++
		temp /= 10
	}

	// Calculate the sum of each digit raised to the power of numDigits
	temp = number
	sum := 0
	for temp > 0 {
		digit := temp % 10
		power := 1
		for i := 0; i < numDigits; i++ {
			power *= digit
		}
		sum += power
		temp /= 10
	}

	// Check if the number is an Armstrong number
	if sum == number {
		fmt.Printf("%d is an Armstrong number.\n", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.\n", number)
	}
}
