package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

const conferenceTickets = 50
var conferenceName = "Go_Conference"
var remainingTickets = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName        string
	lastName         string
	email            string
	numberOfTickets  int
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers(conferenceName, remainingTickets, conferenceTickets)


	firstName, lastName, email, userTickets := getuserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.DataValidation(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(firstName, lastName, email, userTickets)
		wg.Add(1)
		go sendTicket()
		var firstNames = getFirstNames()
		fmt.Println("First names of bookings are", firstNames)

		if remainingTickets == 0 {
			fmt.Println("We are sold out!")
		}
	} else {
		if !isValidName {
			fmt.Println("Please enter a valid name")
		}
		if !isValidEmail {
			fmt.Println("Please enter a valid email")
		}
		if !isValidTicketNumber {
			fmt.Println("Please enter a valid ticket number")
		}
	}
	wg.Wait()
}


func greetUsers(confName string, confTickets int, confRemainingTickets int) {
	fmt.Println("Welcome to our Go program!", confName, "coming up in 2024!")
	fmt.Println("We have", confRemainingTickets, "tickets remaining out of", confTickets, "tickets!")
	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getuserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets int) {
	remainingTickets = remainingTickets - userTickets
	fmt.Println("We have", remainingTickets, "tickets remaining out of", conferenceTickets, "tickets!")

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Println(bookings)
	fmt.Printf("type of bookings %T \n", bookings)

	fmt.Println("Thank you for purchasing", userTickets, "tickets,", firstName, lastName, "you will receive an confirmation email!", email)
}


func sendTicket(){
	time.Sleep(10 * time.Second)
	fmt.Println("Sending ticket to user", bookings[0].firstName, bookings[0].lastName, "at", bookings[0].email, "for", bookings[0].numberOfTickets, "tickets!")
	wg.Done()
}