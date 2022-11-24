/*
 * Write a program to sort an array of integers. The program should
 * partition the array into 4 parts, each of which is sorted by a
 * different goroutine. Each partition should be of approximately
 * equal size. Then the main goroutine should merge the 4 sorted
 * subarrays into one large sorted array.
 *
 * The program should prompt the user to input a series of
 * integers. Each goroutine which sorts ¼ of the array should print
 * the subarray that it will sort. When sorting is complete, the main
 * goroutine should print the entire sorted list.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// var n []int

	// n := []int{6, 8, -1, 3, 10, 1, 45, -12, 6, 9, 2, 5}

	if n, err := ReadValues(); err != nil {
		fmt.Printf("Error reading text from console: %v\n", err)
	} else {
		s := len(n) / 4
		s1 := n[:s]
		s2 := n[s : s*2]
		s3 := n[s*2 : s*3]
		s4 := n[s*3:]

		fmt.Printf("Arreglo desordenado: %v\n", n)

		wg.Add(4)
		go func() {
			sortSegment(s1)
			wg.Done()
		}()
		go func() {
			sortSegment(s2)
			wg.Done()
		}()
		go func() {
			sortSegment(s3)
			wg.Done()
		}()
		go func() {
			sortSegment(s4)
			wg.Done()
		}()
		wg.Wait()

		fmt.Printf("Arreglo ordenado   : %v\n", sortMainArray(s1, s2, s3, s4))
	}
}

func sortSegment(segmento []int) []int {
	sort.Ints(segmento)
	return segmento
}

func sortMainArray(a1, a2, a3, a4 []int) []int {
	var newArray []int

	newArray = append(a1, a2...)
	newArray = append(newArray, a3...)
	newArray = append(newArray, a4...)
	sort.Ints(newArray)
	return newArray
}

func ReadValues() (numbers []int, err error) {
	fmt.Println("Please input numbers(separate with space): ")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine() // ReadLine returns
	ns := strings.Split(string(a), " ")
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return
}
