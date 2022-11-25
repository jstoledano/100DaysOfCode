package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	var (
		argv1 string
		argv2 string
		argv3 string

		zoo map[string]Animal
	)

	zoo = make(map[string]Animal)

	fmt.Printf("Write a command: newanimal or query\n")

	for true {
		fmt.Printf("-> ")
		fmt.Scanf("%s %s %s\n", &argv1, &argv2, &argv3)

		if strings.ToLower(argv1) != "newanimal" && strings.ToLower(argv1) != "query" {
			usage()
		} else if strings.ToLower(argv1) == "newanimal" {
			switch strings.ToLower(argv3) {
			case "cow":
				zoo[argv2] = Cow{name: argv2}
				fmt.Printf("Created it!\n")
			case "bird":
				zoo[argv2] = Bird{name: argv2}
				fmt.Printf("Created it!\n")
			case "snake":
				zoo[argv2] = Snake{name: argv2}
				fmt.Printf("Created it!\n")
			default:
				fmt.Errorf("There is not an animal %v", argv2)
				usage()
			}
		} else if strings.ToLower(argv1) == "query" {
			if _, ok := zoo[argv2]; ok {
				switch strings.ToLower(argv3) {
				case "eat":
					zoo[argv2].Eat()
				case "move":
					zoo[argv2].Move()
				case "speak":
					zoo[argv2].Speak()
				default:
					fmt.Errorf("There is not such verb %v", argv3)
					usage()
				}
			} else {
				fmt.Printf("There is not such animal %v\n", argv2)
			}
		}
	}
}

type Cow struct{ name string }

func (c Cow) Eat()   { fmt.Printf("The cow %s eats grass\n", c.name) }
func (c Cow) Move()  { fmt.Printf("The cow %s locomotion is walk\n", c.name) }
func (c Cow) Speak() { fmt.Printf("The sound of the cow %s is moo\n", c.name) }

type Bird struct{ name string }

func (b Bird) Eat()   { fmt.Printf("The bird %s eats worms\n", b.name) }
func (b Bird) Move()  { fmt.Printf("The bird %v locomotion is fly\n", b.name) }
func (b Bird) Speak() { fmt.Printf("The bird %s sound is peep\n", b.name) }

type Snake struct{ name string }

func (s Snake) Eat()   { fmt.Printf("The snake %s eats mice\n", s.name) }
func (s Snake) Move()  { fmt.Printf("The snake %s locomotion is slither\n", s.name) }
func (s Snake) Speak() { fmt.Printf("The snake %s sound is hsss", s.name) }

func usage() {
	fmt.Println("Usage 1: newanimal <animal_name> <animal_type>")
	fmt.Println("         where animal_name is a string")
	fmt.Println("         where animal_type is one of cow, bird or snake")
	fmt.Println("Usage 2: query <animal_name> <request>")
	fmt.Println("         where animal_name is a string")
	fmt.Println("         where request is one of eat move or speak")
}
