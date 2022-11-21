package main

import "fmt"

type Animal struct {
	food,
	move,
	sound string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.move
}

func (a Animal) Speak() string {
	return a.sound
}

func main() {
	var (
		zoo    map[string]Animal
		animal string
		verb   string
	)

	zoo = make(map[string]Animal)
	zoo["cow"] = Animal{"grass", "walk", "moo"}
	zoo["bird"] = Animal{"worms", "fly", "peep"}
	zoo["snake"] = Animal{"mice", "slither", "hsss"}

	fmt.Printf("Write an animal (cow, bird o snake) and a verb (eat, speak or move):\n")

	for true {
		fmt.Printf("-> ")
		if _, err := fmt.Scanf("%s %s\n", &animal, &verb); err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Printf("The %v", animal)
		switch verb {
		case "eat":
			fmt.Printf(" eats %v\n", zoo[animal].Eat())
		case "move":
			fmt.Printf(" move is %v\n", zoo[animal].Move())
		case "speak":
			fmt.Printf(" sound is %v\n", zoo[animal].Speak())
		default:
			fmt.Printf(", I have not that information or animal.\n")
		}

	}
}
