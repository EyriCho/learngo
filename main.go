package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// lesson 6
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
	third := 1.0 / 3
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%06.2f\n", third)

	var balance float32

	for balance < 20.00 {
		var coin = rand.Intn(3);
		var amount float32;
		switch coin{
		case 0:
			amount = 0.05
		case 1:
			amount = 0.1
		default:
			amount = 0.25
		}
		balance += amount
		fmt.Printf("After deposit $%6.2f, current balance is %.2f\n", amount, balance)
	}
}