package main

import "fmt"

func main() {
	a := 0
	b := 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 1; i <= 10; i++ {
		sum := a + b
		fmt.Println(sum)
		a = b
		b = sum
	}

}
