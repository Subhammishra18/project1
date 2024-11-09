package main

import (
	"fmt"
)

func main() {
	var ans int
	for {
		var op string
		fmt.Print("enter the operation(+ ,-, *,% ,x to exit): ")
		fmt.Scan(&op)
		if op == "x" {
			fmt.Println("exiting th program")
			break

		}
		var num1, num2 int
		fmt.Println("enter the first number:")
		fmt.Scan(&num1)
		fmt.Println("enter the second number:")
		fmt.Scan(&num2)
		switch op {
		case "+":
			ans = num1 + num2
		case "-":
			ans = num1 - num2
		case "*":
			ans = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("error :division by 0 is not allowed")
				continue

			}
			ans = num1 / num2
		case "%":
			ans = num1 % num2
		default:
			fmt.Println("invlaid ")

		}
		fmt.Println(" answer of your operation ->", ans)

	}

}
