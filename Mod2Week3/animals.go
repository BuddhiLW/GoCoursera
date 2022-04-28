package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

type Animals struct {
	Animals []Animal
}

func (a Animal) Eat() {
	fmt.Println("Food eaten:", a.Food)
}

func (a Animal) Walk() {
	fmt.Println("Locomotion method:", a.Locomotion)
}

func (a Animal) Speak() {
	fmt.Println("Spoken sound:", a.Noise)
}

type DoubleStrLists struct {
	list1 []string
	list2 []string
}

func (lists DoubleStrLists) PossibleStrings(input string) (string, string) {
	var animal string
	var behavior string

	for _, string1 := range lists.list1 {
		for _, string2 := range lists.list2 {
			var concatStr = []string{string1, string2}
			if strings.Join(concatStr, " ") == input {
				fmt.Println("It's in the possible list, in which", "animal:", concatStr[0], "behavior:", concatStr[1])
				animal, behavior = concatStr[0], concatStr[1]
			}
		}
	}
	return animal, behavior
}

// func (lists DoubleStrLists) PossibleStringsPrint() {
// 	for _, string1 := range lists.list1 {
// 		for _, string2 := range lists.list2 {
// 			var concatStr = []string{string1, string2}
// 			fmt.Println(strings.Join(concatStr, " "))
// 		}
// 	}
// }

func main() {
	cow := Animal{Food: "grass", Locomotion: "walk", Noise: "moo"}
	bird := Animal{Food: "worms", Locomotion: "fly", Noise: "peep"}
	snake := Animal{Food: "mice", Locomotion: "slither", Noise: "hsss"}

	fmt.Println("Use key-words in lower-case. \nE.g., cow eat, bird move, snake speak\n If you want to exit, press 'X' and Enter")

	for {
		var input string
		fmt.Println(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input = scanner.Text()
		fmt.Println("captured:", input)
		if input == "X" {
			break
		} else {
			var strings DoubleStrLists
			strings = DoubleStrLists{list1: []string{"cow", "bird", "snake"}, list2: []string{"eat", "move", "speak"}}

			animal, behavior := strings.PossibleStrings(input)

			switch animal {
			case "cow":
				switch behavior {
				// Eat Walk Speak
				case "eat":
					cow.Eat()
				case "move":
					cow.Walk()
				case "speak":
					cow.Speak()
				}

			case "bird":
				switch behavior {
				// Eat Walk Speak
				case "eat":
					bird.Eat()
				case "move":
					bird.Walk()
				case "speak":
					bird.Speak()
				}
			case "snake":
				switch behavior {
				// Eat Walk Speak
				case "eat":
					snake.Eat()
				case "move":
					snake.Walk()
				case "speak":
					snake.Speak()
				}
			}
		}
	}
}
