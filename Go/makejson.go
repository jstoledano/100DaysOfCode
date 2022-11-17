package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var name, address string
	fmt.Printf("Write your name: ")
	fmt.Scan(&name)
	fmt.Printf("Write yout address: ")
	reader := bufio.NewReader(os.Stdin)
	address := bufio.NewScanner(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nYour name is %v and you live in %v\n", name, address)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	j, _ := json.Marshal(m)
	fmt.Println(m)
	fmt.Println(string(j))
}
