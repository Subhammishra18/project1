package main

import "fmt"

func main() {
	arr := []int{9, 65, 2, 57, 54, 13, 67, 99, 1234, 4, 2, 1}
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
