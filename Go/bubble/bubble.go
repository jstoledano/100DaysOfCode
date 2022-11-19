package main

import "fmt"

func main() {
	var (
		numbers []int
		cap     int = 10
	)

	fmt.Printf("=== Write 10 integers ===\n")
	for i := 0; i < cap; i++ {
		var number int
		fmt.Printf("Integer #%d: ", i+1)
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}

	fmt.Printf("Unordered slice\n")
	fmt.Printf("%v\n", numbers)

	bubbleSort(numbers)

	fmt.Printf("Ordered Slice\n")
	fmt.Printf("%v\n", numbers)
}

func swap(slice []int, index int) bool {
	slice[index+1], slice[index] = slice[index], slice[index+1]
	return true
}

func bubbleSort(data []int) {
	var (
		n      = len(data)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if data[i] > data[i+1] {
				swapped = swap(data, i)
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
