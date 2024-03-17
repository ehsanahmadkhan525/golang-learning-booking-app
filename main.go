package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var conferneceName = "Go_Conference"
	const conferneceTickets = 50
	var remaingTickets = 50
	fmt.Println("Welcome to our Go program!", conferneceName, "coming up in 2024!")
	fmt.Println("We have", remaingTickets, "tickets remaining out of", conferneceTickets, "tickets!")
	fmt.Println("Get your tickets now!")
    
	var bookings []string // Move the declaration outside the loop
	
	for{
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
	
	
		remaingTickets = remaingTickets - userTickets
		fmt.Println("We have", remaingTickets, "tickets remaining out of", conferneceTickets, "tickets!")
	
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Println(bookings)
		fmt.Printf("type of bookings %T \n", bookings)
	
		fmt.Println("Thank you for purchasing", userTickets, "tickets,", firstName, lastName, "you will receive an confirmation email!", email)
	
		firstNames :=[]string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("first names of bookings are %v \n ", firstNames)
	}
}