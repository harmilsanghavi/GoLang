package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var confreshName = "hello"

const confreshTicket = 50

var remainingTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noofTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("hello")
	fmt.Printf("maths %v \n", rand.Intn(10))
	greetUser()

	userName, userLastName, userEmail, userTicket := userInput()

	isValidName, isValidEmail, isValidTicketNumber := Validation(userName, userLastName, userEmail, userTicket)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTicket, userName, userLastName, userEmail)
		wg.Add(1)
		go sendTicket(userTicket, userName, userLastName, userEmail)

		s := printFirstName()
		fmt.Printf("the first names of booking %v\n", s)

		if remainingTicket == 0 {
			fmt.Println("ticket are all sold")
			// break
		}
	} else {
		if !isValidName {
			fmt.Printf("your first name or last name is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("your email doesn't contain @\n")
		}
		if !isValidTicketNumber {
			fmt.Printf("numvber of tickets you enter is invalid\n")
		}
	}
	wg.Wait()
}
func greetUser() {
	fmt.Printf("type of confreshName %T\n", confreshName)
	fmt.Println("welcome", confreshName)
	fmt.Println("welcome", confreshTicket)
	fmt.Println("welcome", remainingTicket)
}
func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	//fmt.Printf("the first names of booking %v\n", firstNames)
}
func userInput() (string, string, string, uint) {
	var userName string
	var userLastName string
	var userEmail string
	var userTicket uint

	fmt.Printf("enter your first name\n")
	fmt.Scan(&userName)

	fmt.Printf("enter your last name\n")
	fmt.Scan(&userLastName)

	fmt.Printf("enter your email name\n")
	fmt.Scan(&userEmail)

	fmt.Printf("enter your ticket name\n")
	fmt.Scan(&userTicket)

	return userName, userLastName, userEmail, userTicket
}
func bookTicket(userTicket uint, userName string, userLastName string, userEmail string) {
	remainingTicket = remainingTicket - userTicket

	var userData = UserData{
		firstName:   userName,
		lastName:    userLastName,
		email:       userEmail,
		noofTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of Bookings is %v\n", bookings)

	fmt.Printf("user %v %v booked %v tickets on %v these email\n", userName, userLastName, userTicket, userEmail)
	fmt.Printf("%v tickets is remaing\n", remainingTicket)
}
func sendTicket(userTicket uint, userName string, userLastName string, userEmail string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTicket, userName, userLastName)
	fmt.Println("########")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n", ticket, userEmail)
	fmt.Println("########")
	wg.Done()
}
