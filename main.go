package main

import (
	"fmt"
	"strings"
)

func main() {
	// lesson 3
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
	if command == "walk outside" {
		fmt.Println("You head further up the mountain")
	} else if command == "walk inside" {
		fmt.Println("You head further deep in the cavern")
	} else {
		fmt.Println("You have completely losed")
	}
	var room = "lake"
	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern")
	case "lake":
		fmt.Println("The ice seems solid enough")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold")
	}

	var num = 10
	for num > 0 {
		fmt.Printf("%-4v", num)
		num--
	}
}