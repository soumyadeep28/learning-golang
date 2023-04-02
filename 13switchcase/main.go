package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the lession of switchcase ")
	//random
	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1
	fmt.Println("dice number :", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough //this is the case where it will go down to the next case aswell
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("What is that ")
	}

}
