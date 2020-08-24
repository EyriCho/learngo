package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%-16v %-4v %-11v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
	count := 39
	for count > 0 {
		fmt.Print("=")
		count--
	}
	fmt.Print("\n")
	for count = 0; count < 10; count++ {
		spaceline := "Virgin Galactic"
		days := 23
		tripType := "Round-trip"
		price := 96
		switch count {
		case 0, 1, 5, 6:
			spaceline = "Virgin Galactic"
		case 2, 9:
			spaceline = "SpaceX"
		case 3, 4, 7, 8:
			spaceline = "Space Adventure"
		}
		switch count {
		case 1, 2, 4, 7:
			tripType = "One-way"
		}
		switch count {
		case 1:
			days = 39
			price = 37
		case 2:
			days = 31
			price = 41
		case 3:
			days = 22
			price = 100
		case 4:
			days = 22
			price = 50
		case 5:
			days = 30
			price = 84
		case 6:
			days = 24
			price = 94
		case 7:
			days = 27
			price = 44
		case 8:
			days = 28
			price = 86
		case 9:
			days = 41
			price = 72
		}
		fmt.Printf("%-16v %4v %-11v $%4v\n", spaceline, days, tripType, price)
	}
}