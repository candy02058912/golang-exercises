package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func selectAnimal() string {
	fmt.Println("Animal type: (cow, bird, snake)")
	var input string
	fmt.Scan(&input)
	return input
}

func selectAction() string {
	fmt.Println("Action type: (eat, move, speak)")
	var input string
	fmt.Scan(&input)
	return input
}

func performAction(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("Action type %v does not exist\n", action)
	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	animal := selectAnimal()
	switch animal {
	case "cow":
		performAction(cow, selectAction())
	case "bird":
		performAction(bird, selectAction())
	case "snake":
		performAction(snake, selectAction())
	default:
		fmt.Printf("Animal type %v does not exist.\n", animal)
	}
}
