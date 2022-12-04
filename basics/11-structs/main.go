package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Structs")
	fmt.Println("")

	myuser := User{"saurav", "sauravrana646@gmail.com", false, 22}

	fmt.Println("User created is ",myuser)
	fmt.Printf("User create is %+v\n",myuser)
	fmt.Println("User's Name is ",myuser.Name)

}

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        uint8
}
