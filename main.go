package main

import (
	"fmt"
	"math/big"
)

func main() {
	// lesson 8

	// big package
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed")

	// constant
	const lightSpeedConst = 299792
	const secondsPerDayConst = 86400
	const daysPerYearConst = 365
	fmt.Println("Canis Major Dwarf is", 236000000000000000 / lightSpeedConst / secondsPerDayConst / daysPerYearConst, "light years away.")
}