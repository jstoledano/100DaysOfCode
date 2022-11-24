package main

import (
	"fmt"
	"time"
)

// This is a simple function where occurs the race condition.  It
// takes as a parameter the pointer to a variable. Inside n is
// multiplied by two and print the result.
//
// But this function takes one second to finish, in that time the
// variable is used again, but use its original value because the
// process has not been completed yet.

func twoTimes(n *int32) int32 {
	time.Sleep(time.Second * 1)
	*n *= 2
	return *n
}

func main() {
	var n int32 = 3
	fmt.Printf("Before the Goroutine n=%v\n", n)
	go func() {
		datos := twoTimes(&n)
		fmt.Printf("Inside the Gorutine n=%v\n", datos)
	}()
	fmt.Printf("After the Goroutine n=%v\n", n)
	time.Sleep(time.Second * 3)
	fmt.Printf("After the sleep n=%v\n", n)
}
