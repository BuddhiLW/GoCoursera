package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	// cow := Animal{Food: "grass", Locomotion: "walk", Noise: "moo"}
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

// func (as Animals) anyAnimal() string {

// }

func main() {
	cow := Animal{Food: "grass", Locomotion: "walk", Noise: "moo"}
	bird := Animal{Food: "worms", Locomotion: "fly", Noise: "peep"}
	snake := Animal{Food: "mice", Locomotion: "slither", Noise: "hsss"}
	animals := Animals{Animals: []Animal{cow, bird, snake}}

	// var animals []Animal
	// animals = make([]Animal, 3)
	// animals[0] = cow
	// animals[1] = bird

	// var animals Animals
	// animals = make(animals2, 0, 3)

	// var input string
	// if input != "X" {
	// 	fmt.Println(">")
	// 	fmt.Scan(&input)
	// }
	v := reflect.ValueOf(cow)
	// V := reflect.ValueOf(animals)
	typeOfv := v.Type()
	// typeOfV := reflect.ValueOf(animals)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfv.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println(animals)

	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("Field: %s\tValue: %v\n", typeOfV.Field(i).Name, V.Field(i).Interface())
	// }
}
