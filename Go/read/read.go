package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
  Write a program which reads information from a file and represents
  it in a slice of structs. Assume that there is a text file which
  contains a series of names. Each line of the text file has a first
  name and a last name, in that order, separated by a single space on
  the line.

  Your program will define a name struct which has two fields, fname
  for the first name, and lname for the last name. Each field will be
  a string of size 20 (characters).

  Your program should prompt the user for the name of the text
  file. Your program will successively read each line of the text file
  and create a struct which contains the first and last names found in
  the file. Each struct created will be added to a slice, and after
  all lines have been read from the file, your program will have a
  slice containing one struct for each line in the file. After reading
  all lines from the file, your program should iterate through your
  slice of structs and print the first and last names found in each
  struct.
*/

func main() {
	// Constants
	const (
		MaxNameLength = 20
	)

	// Variables
	var filename string

	type name struct {
		fname string
		lname string
	}

	nameSlice := make([]name, 0)

	// Prompt for filename
	fmt.Printf("Enter filename: ")
	fmt.Scan(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var firstName, lastName string
		fmt.Sscanf(scanner.Text(), "%s %s", &firstName, &lastName)
		nameSlice = append(nameSlice, name{firstName, lastName})
	}

	// Print the slice
	fmt.Printf(" # |      First Name      |      Last Name\n")
	fmt.Printf("---+----------------------+----------------------\n")

	for i, v := range nameSlice {
		fmt.Printf(" %v | %-20s | %-20s\n",
			i+1, cleanLength(v.fname, MaxNameLength),
			cleanLength(v.lname, MaxNameLength))
	}
}

func cleanLength(s string, l int) string {
	if len(s) > l {
		return s[:l]
	}
	return s
}
