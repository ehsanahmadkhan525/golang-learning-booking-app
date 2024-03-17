package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferneceName = "Go_Conference"
	const conferneceTickets = 50
	var remaingTickets = 50
	greetUsers(conferneceName, remaingTickets, conferneceTickets)


	var bookings []string // Move the declaration outside the loop

	for {

		firstName, lastName, email , userTickets :=  getuserInput()


        isValidName, isValidEmail, isValidTicketNumber := dataValidation(firstName, lastName, email, userTickets, remaingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remaingTickets, conferneceTickets, firstName, lastName, email, userTickets, bookings)

            var firstNames = getFirstNames(bookings)
			fmt.Println("First names of bookings are", firstNames)

			if remaingTickets == 0 {
				fmt.Println("We are sold out!")

		} 
	}else {
		if !isValidName{
			fmt.Println("Please enter a valid name")
		}
		if !isValidEmail{
			fmt.Println("Please enter a valid email")
		}
		if !isValidTicketNumber{
			fmt.Println("Please enter a valid ticket number")}

	}
}
}



func greetUsers(confName string, confTickets int, confRemainingTickets int){
	fmt.Println("Welcome to our Go program!", confName, "coming up in 2024!")
	fmt.Println("We have", confRemainingTickets, "tickets remaining out of", confTickets, "tickets!")
	fmt.Println("Get your tickets now!")
}


func getFirstNames(bookings []string) []string{

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("first names of bookings are %v \n ", firstNames)

	return firstNames
}

func dataValidation(firstName string, lastName string, email string, userTickets int, remaingTickets int) (bool, bool, bool){

	var isValidName = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}


func getuserInput()(string, string, string, int){
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


func bookTicket(remaingTickets int, conferneceTickets int, firstName string, lastName string, email string, userTickets int, bookings []string){
	remaingTickets = remaingTickets - userTickets
	fmt.Println("We have", remaingTickets, "tickets remaining out of", conferneceTickets, "tickets!")

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println(bookings)
	fmt.Printf("type of bookings %T \n", bookings)

	fmt.Println("Thank you for purchasing", userTickets, "tickets,", firstName, lastName, "you will receive an confirmation email!", email)
}