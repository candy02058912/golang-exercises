package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func readInput() []string {
	fmt.Println("Commands available:")
	fmt.Println("newanimal <name> <type>(cow, bird, snake)")
	fmt.Println("query <name> <info>(eat, move, speak)")
	fmt.Print("> ")
	input := make([]string, 3)
	for i := range input {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			fmt.Errorf("error with input %v: %v", input[:i], err)
		}
	}
	return input
}

func createAnimal(animalType, name string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{name}, nil
	case "bird":
		return Bird{name}, nil
	case "snake":
		return Snake{name}, nil
	default:
		return nil, fmt.Errorf("Animal type doesn't exist, only cow, bird, snake are available.")
	}
}

func requestInfo(animal Animal, info string) {
	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Info type doesn't exist, only eat, move, speak are available.")
	}
}

func main() {
	animals := map[string]Animal{}
	for {
		input := readInput()
		command, name := input[0], input[1]
		switch command {
		case "newanimal":
			animalType := input[2]
			_, ok := animals[name]
			if ok {
				fmt.Printf("%v already exists\n", name)
			} else {
				animal, err := createAnimal(animalType, name)
				if err != nil {
					fmt.Println(err)
				}
				animals[name] = animal
				fmt.Printf("%v is created!\n", name)
			}
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Printf("%v does not exist yet\n", name)
			} else {
				info := input[2]
				requestInfo(animal, info)

			}
		}
	}
}
