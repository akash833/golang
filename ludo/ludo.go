package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the ludo game")

	number := rand.Intn(6) + 1

	switch number {
	case 1:
		fmt.Println("Dice value is open and you can move")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 3")
	case 5:
		fmt.Println("Dice value is 3")
	case 6:
		fmt.Println("Dice value is 3")
	default:
		fmt.Println("Byy bbyy")
	}
}
