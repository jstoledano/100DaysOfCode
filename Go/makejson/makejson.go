package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Variables
	var name, address string

	// Prompt for name and address
	fmt.Printf("Write your name: ")
	fmt.Scan(&name)
	fmt.Printf("Write yout address: ")
	reader := bufio.NewReader(os.Stdin)
	address, _ = reader.ReadString(('\n'))

	// Making a map
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	// Create the json
	j, _ := json.Marshal(m)
	fmt.Println(string(j))
}
