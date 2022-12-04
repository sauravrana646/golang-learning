package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Arrays")
	fmt.Println("")

	var fruitList [7]string
	var ptr *[7]string = &fruitList

	fmt.Println("Length of array before is ", len(fruitList))

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[6] = "Mango"

	fmt.Println("Length of slice after is ", len(fruitList))

	fmt.Println("The Fruits List is ", fruitList)
	fmt.Println("The Address of Fruits List is ", ptr)
}
