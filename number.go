package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

}
