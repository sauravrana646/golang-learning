package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang pointers")
	fmt.Println("")

	// Initialize an empty pointer
	// var ptr *int

	num := 23

	ptr := &num

	fmt.Printf("The value of ptr is : %v\n", ptr)
	fmt.Printf("The value inside ptr address is : %v\n\n", *ptr)
	
	*ptr += 5
	
	fmt.Printf("The value of ptr is : %v\n", ptr)
	fmt.Printf("The value inside ptr address is : %v\n\n", *ptr)

}