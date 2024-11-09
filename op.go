package main

import (
	"fmt"
)

func main() {
	var ans int
	for {
		// Get the operator
		var op string
		fmt.Print("Enter operation (+, -, *, /, %, x to exit): ")
		fmt.Scan(&op)

		// If user wants to exit
		if op == "x" {
			fmt.Println("Exiting the program.")
			break
		}

		// Get the two numbers
		var num1, num2 int
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		// Perform the operation
		switch op {
		case "+":
			ans = num1 + num2
		case "-":
			ans = num1 - num2
		case "*":
			ans = num1 * num2
		case "/":
			
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				continue 
			}
			ans = num1 / num2
		case "%":
			ans = num1 % num2
		default:
			fmt.Println("Invalid operator.")
			continue 
		}

		
		fmt.Printf("Result: %d\n", ans)
	}
}
