package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"

)

const conferenceTickets = 50
var conferenceName = "Go_Conference"

var remainingTickets = 50
// var bookings []string 
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers(conferenceName, remainingTickets, conferenceTickets)

	for {
		firstName, lastName, email , userTickets :=  getuserInput()

        isValidName, isValidEmail, isValidTicketNumber := helper.DataValidation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = bookTicket(remainingTickets, conferenceTickets, firstName, lastName, email, userTickets, bookings)
            var firstNames = getFirstNames(bookings)
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
	}
}

func greetUsers(confName string, confTickets int, confRemainingTickets int) {
	fmt.Println("Welcome to our Go program!", confName, "coming up in 2024!")
	fmt.Println("We have", confRemainingTickets, "tickets remaining out of", confTickets, "tickets!")
	fmt.Println("Get your tickets now!")
}



func getFirstNames(bookings []map[string]string) []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking["firstName"])
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

func bookTicket(remaingTickets int, conferneceTickets int, firstName string, lastName string, email string, userTickets int, bookings []map[string]string) []map[string]string{
	remaingTickets = remaingTickets - userTickets
	fmt.Println("We have", remaingTickets, "tickets remaining out of", conferneceTickets, "tickets!")


	var userData =make (map [string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData ["numberOfTickets"]= strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Println(bookings)
	fmt.Printf("type of bookings %T \n", bookings)

	fmt.Println("Thank you for purchasing", userTickets, "tickets,", firstName, lastName, "you will receive an confirmation email!", email)
	return bookings
}
