package main

import "fmt"

func main() {
	n := 1234567
	ans := 0
	for n > 0 {
		num := n % 10
		n = n / 10
		ans = ans*10 + num
	}
	fmt.Println(ans)

}
