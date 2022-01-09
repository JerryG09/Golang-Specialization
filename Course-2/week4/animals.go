package main

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

// Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (ani animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var subjectAnimal animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		_, err := fmt.Scan(&command, &requestAni, &requestType)
		if err != nil {
			fmt.Println("Error Occurred while reading input: ", err)
		}
		if command == "query" {
			subjectAnimal = animalMap[requestAni]
			switch requestType {
			case "eat":
				subjectAnimal.Eat()
			case "move":
				subjectAnimal.Move()
			case "speak":
				subjectAnimal.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}
