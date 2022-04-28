package main

import (
	"fmt"
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
	fmt.Println("Food eaten", a.Food)
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

func (lists DoubleStrLists) PossibleStrings(input string) {
	for _, string1 := range lists.list1 {
		for _, string2 := range lists.list2 {
			var concatStr = []string{string1, string2}
			if strings.Join(concatStr, " ") == input {
				fmt.Println(concatStr[0], concatStr[1])
			} else {
				fmt.Println("Either the animal-behavior input is incorrect,\n or the animal-behavior is not listed")
			}
		}
	}
}

func main() {
	// cow := Animal{Food: "grass", Locomotion: "walk", Noise: "moo"}
	// bird := Animal{Food: "worms", Locomotion: "fly", Noise: "peep"}
	// snake := Animal{Food: "mice", Locomotion: "slither", Noise: "hsss"}

	fmt.Println("Use key-words in lower-case. E.g., cow eat, bird move, snake speak")

	var input string
	if input != "X" {
		fmt.Println(">")
		fmt.Scan(&input)

		var strings DoubleStrLists
		strings = DoubleStrLists{list1: []string{"cow", "bird", "snake"}, list2: []string{"eat", "move", "speak"}}

		strings.PossibleStrings(input)
		// if strings.PossibleStrings(input) {

		// }
	}

	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Printf("Field: %s\tValue: %v\n", typeOfv.Field(i).Name, v.Field(i).Interface())
	// }

	// fmt.Println(animals)
}
