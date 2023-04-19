package main

import "strings"

func Validation(userName string, userLastName string, userEmail string, userTicket uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
