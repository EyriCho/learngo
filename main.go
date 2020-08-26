package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// lesson 7
	// introduce value type
	var year = 2018
	day := 365.2425
	word := "Hello world;"
	escaped := false
	fmt.Printf("Type %T for %v\n", year, year)
	fmt.Printf("Type %T for %v\n", day, day)
	fmt.Printf("Type %T for %v\n", word, word)
	fmt.Printf("Type %T for %v\n", escaped, escaped)

	// hexadecimal
	var red, green, yellow = 0x00, 0x8d, 0xd5
	fmt.Printf("color: #%02x%02x%02x\n", red, green, yellow)

	// integer wrap around
	var max, min uint8 = 255, 0
	fmt.Printf("%08b\n%08b\n", max, min)
	max++
	min--
	fmt.Printf("%08b\n%08b\n", max, min)

	// time wrap around
	future := time.Unix(12622780800, 0)
	fmt.Println(future)

	// piggy bank
	var nickels, dimes, quarters int16 = 0, 0, 0

	var balance float32
	for balance < 20.00 {
		var coin = rand.Intn(3)
		var amount float32
		switch coin{
		case 0:
			amount = 0.05
			nickels++
		case 1:
			amount = 0.1
			dimes++
		default:
			amount = 0.25
			quarters++
		}
		balance = 0.05 * float32(nickels) + 0.1 * float32(dimes) + 0.25 * float32(quarters)
		fmt.Printf("After deposit $%6.2f, current balance is %.2f\n", amount, balance)
	}
}