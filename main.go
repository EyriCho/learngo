package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// lesson 2
	fmt.Println("Lesson 2")
	const (
		distance = 56000000
	)
	var (
		speed = 100800
		randNum = rand.Intn(10) + 1
	)
	fmt.Println("Random number is ", randNum)
	fmt.Println("Take ", distance / speed, "days travel to Mars")
	fmt.Print("My weight on mars is ")
	fmt.Print(149.0 * 0.3782)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")
}