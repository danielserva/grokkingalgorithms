package main

import "fmt"

func binarySearch(lista []int, item int) int {
	low := 0
	high := len(lista) - 1
	for low <= high {
		fmt.Printf("low: %d high: %d\n", low, high)
		half := (low + high) / 2
		fmt.Printf("half: %d\n", half)
		guess := lista[half]
		fmt.Printf("guess: %d\n", guess)
		if guess == item {
			return half
		}
		if guess > item {
			high = half - 1
		} else {
			low = half + 1
		}
	}
	return -1
}

func main() {
	myList := []int{1, 3, 5, 7, 9}
	fmt.Printf("position of item 3: %d \n", binarySearch(myList, 3))
}
