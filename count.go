package main

import "fmt"

func main() {
	n := 1234456765456765
	count := 0

	for n > 0 {
		rem := n % 10
		if rem == 5 {
			count++
		}
		n = n / 10
	}
	fmt.Println(count)
}
