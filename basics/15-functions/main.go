package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Functions")
	fmt.Println("")

	greets()

	sum := adder(2,3)
	fmt.Println("Sum is ", sum)

	prosum,message := proadder(2,5,5,8,3)
	fmt.Println("Pro sum is ", prosum)
	fmt.Println("Message says ", message)
}

func proadder(a ... int) (int,string) {
	sum := 0
	for _ , value := range a {
		sum += value
	}
	return sum, "Hi from proadder"
}

func adder(a int , b int) int {
	return a+b
}

func greets(){
	fmt.Println("Hello There...!!")
}