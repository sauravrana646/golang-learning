package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang Switch Case")
	fmt.Println("")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6)+1

	fmt.Println("Dice Number is ",diceNum)

	switch diceNum {
	case 1:
		fmt.Println("You are allowed to open")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 space")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces and roll dice again")
	default:
		fmt.Println("What was that !!")
		
	}

}
