package main

import (
	"strings"
)

func userInputValidation(firstName string , lastName string , userEmail string , userTicket uint)(bool,bool,bool){

	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicket
}