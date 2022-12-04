package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets uint = 100

var conferenceName string = "Go Conference"   // Way to declare variable
var remainingTicket uint = 100   
var bookings = make([] map[string]string, 0)                // Also a way to to declare variable but doesnot work with constants

func main()  {

	
	
	greetUser()

	

	for remainingTicket > 0 {
		
		firstName,lastName,userEmail,userTicket := getUserInput()

		isValidName , isValidEmail , isValidTicket := userInputValidation(firstName,lastName,userEmail,userTicket)	
		
		if isValidEmail && isValidName && isValidTicket {
			bookTicket(firstName , lastName , userEmail , userTicket)
			
		}

		firstNames := getFirstNames()

		fmt.Printf("The First Name of the bookings are : %v\n\n", firstNames)
		
	}

	fmt.Print("All the ticket have been sold out..!! Come back next year.\n")
}

func greetUser(){
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\nGet you tickets to attend here.\n", conferenceTickets, remainingTicket)
}

func getUserInput()(string, string, string,  uint){
	var firstName string
	var lastName string
	var userEmail string
	var userTicket uint

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your Email Address: ")
	fmt.Scan(&userEmail)
	fmt.Print("Enter Number of Tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, userEmail, userTicket
}

func bookTicket(firstName string , lastName string , userEmail string , userTicket uint ){
	remainingTicket = remainingTicket - userTicket

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["userEmail"] = userEmail
	userData["userTicket"] = strconv.FormatUint(uint64(userTicket),10)

	bookings = append(bookings,userData)

	fmt.Printf("Thank You %v %v for booking %v ticket. You will get confirmation on your email at %v\n",firstName,lastName,userTicket,userEmail)
	fmt.Printf("%v tickets remaining for %v\n" , remainingTicket , conferenceName)
}



func getFirstNames()([]string){
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}