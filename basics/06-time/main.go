package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time module practice")

	presentTime := time.Now()

	fmt.Printf("Unformated time now is : %v\nType of time is : %T\n", presentTime, presentTime)

	fmt.Printf("Formated time is : %v", presentTime.Format("02 Jan 2006 15:04 PM Monday"))
	
	createdTime := time.Date(2000, time.June , 17 , 8 , 30 , 0 ,0 , time.Local)
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("Unformated created time is : %v\nType of time is : %T\n", createdTime, createdTime)

	fmt.Printf("Formated created time is : %v", createdTime.Format("02 Jan 2006 15:04 PM Monday"))
	
}
