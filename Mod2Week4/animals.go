package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Walk()
	Speak()
}

type Cow struct {
	Food       string
	Locomotion string
	Noise      string
}

type Bird struct {
	Food       string
	Locomotion string
	Noise      string
}

type Snake struct {
	Food       string
	Locomotion string
	Noise      string
}

// func (a Animal) Walk() {
// 	switch sh := s.(type) {
// 	case Cow
// 	fmt.Println("Locomotion method:", a.Locomotion)
// 	}
// }

// func (a Animal) Speak() {
// 	fmt.Println("Spoken sound:", a.Noise)
// }

// func (a Animal) Eat() {
// 	fmt.Println("Food eaten:", a.Food)
// }

func (a Cow) Eat() {
	fmt.Println("Food eaten:", a.Food)
}

func (a Cow) Walk() {
	fmt.Println("Locomotion method:", a.Locomotion)
}

func (a Cow) Speak() {
	fmt.Println("Spoken sound:", a.Noise)
}

func (a Bird) Eat() {
	fmt.Println("Food eaten:", a.Food)
}

func (a Bird) Walk() {
	fmt.Println("Locomotion method:", a.Locomotion)
}

func (a Bird) Speak() {
	fmt.Println("Spoken sound:", a.Noise)
}

func (a Snake) Eat() {
	fmt.Println("Food eaten:", a.Food)
}

func (a Snake) Walk() {
	fmt.Println("Locomotion method:", a.Locomotion)
}

func (a Snake) Speak() {
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

func main() {
	cow := Cow{Food: "grass", Locomotion: "walk", Noise: "moo"}
	bird := Bird{Food: "worms", Locomotion: "fly", Noise: "peep"}
	snake := Snake{Food: "mice", Locomotion: "slither", Noise: "hsss"}

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
			var s Animal

			var strings DoubleStrLists
			strings = DoubleStrLists{list1: []string{"cow", "bird", "snake"}, list2: []string{"eat", "move", "speak"}}

			animal, behavior := strings.PossibleStrings(input)

			switch animal {
			case "cow":
				switch behavior {
				case "eat":
					s = cow
					s.Eat()
				case "move":
					s = cow
					s.Walk()
				case "speak":
					s = cow
					s.Speak()
				}

			case "bird":
				switch behavior {
				case "eat":
					s = bird
					s.Eat()
				case "move":
					s = bird
					s.Walk()
				case "speak":
					s = bird
					s.Speak()
				}
			case "snake":
				switch behavior {
				case "eat":
					s = snake
					s.Eat()
				case "move":
					s = snake
					s.Walk()
				case "speak":
					s = snake
					s.Speak()
				}
			}
		}
	}
}
