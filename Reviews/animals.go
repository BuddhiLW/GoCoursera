package main

import (
	"fmt"
	"log"
)

type Animals struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animals) InitAnimal(animalName string) {
	switch animalName {
	case "cow":
		animal.food = "grass"
		animal.locomotion = "walk"
		animal.noise = "moo"
	case "bird":
		animal.food = "worms"
		animal.locomotion = "fly"
		animal.noise = "peep"
	case "snake":
		animal.food = "mice"
		animal.locomotion = "slither"
		animal.noise = "hsss"
	}
}

func (animal Animals) Eat() {
	fmt.Println(animal.food)
}

func (animal Animals) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animals) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	inputAnimals()
}

func inputAnimals() {
	fmt.Print(">")
	var animalName, animalTask string
	fmt.Scan(&animalName, &animalTask)
	var cow, bird, snake *Animals
	var currentAnimal **Animals
	var isCowInit, isBirdInit, isSnakeInit bool
	switch animalName {
	case "cow":
		cow = InitAnimal(animalName, &currentAnimal, &isCowInit, cow)
	case "bird":
		bird = InitAnimal(animalName, &currentAnimal, &isBirdInit, bird)
	case "snake":
		snake = InitAnimal(animalName, &currentAnimal, &isSnakeInit, snake)
	default:
		log.Fatal("Exiting, invalid animal")
	}
	for {
		switch animalTask {
		case "eat":
			(*(*currentAnimal)).Eat()
		case "move":
			(*(*currentAnimal)).Move()
		case "speak":
			(*(*currentAnimal)).Speak()
		default:
			log.Fatal("Exiting, invalid task")
		}
		fmt.Print(">")
		fmt.Scan(&animalName, &animalTask)

		switch animalName {
		case "cow":
			if !isCowInit {
				fmt.Println("Here")
				cow = InitAnimal(animalName, &currentAnimal, &isCowInit, cow)
			}
			currentAnimal = &cow
		case "bird":
			if !isBirdInit {
				bird = InitAnimal(animalName, &currentAnimal, &isBirdInit, bird)
			}
			currentAnimal = &bird
		case "snake":
			if !isSnakeInit {
				snake = InitAnimal(animalName, &currentAnimal, &isSnakeInit, snake)
			}
			currentAnimal = &snake
		default:
			log.Fatal("Exiting, invalid animal")
		}

	}
}
func InitAnimal(name string, currentAnimalPointerUpdater ***Animals, animalSetFlag *bool, currAnimal *Animals) *Animals {
	var animal Animals
	animal.InitAnimal(name)
	*animalSetFlag = true
	*currentAnimalPointerUpdater = &currAnimal
	(*(*currentAnimalPointerUpdater)) = &animal
	return &animal
}
