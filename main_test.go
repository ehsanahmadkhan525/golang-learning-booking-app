package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	// Test case 1: Purchase 2 tickets
	t.Run("Purchase 2 tickets", func(t *testing.T) {
		// Set up
		expectedRemainingTickets := 48
		expectedBookings := []string{"John Doe"}

		// Execute
		bookings, remainingTickets := purchaseTickets("John", "Doe", "john.doe@example.com", 2)

		// Verify
		if remainingTickets != expectedRemainingTickets {
			t.Errorf("Expected remaining tickets: %d, got: %d", expectedRemainingTickets, remainingTickets)
		}
		if !stringSliceEqual(bookings, expectedBookings) {
			t.Errorf("Expected bookings: %v, got: %v", expectedBookings, bookings)
		}
	})

	// Test case 2: Purchase 0 tickets
	t.Run("Purchase 0 tickets", func(t *testing.T) {
		// Set up
		expectedRemainingTickets := 50
		expectedBookings := []string{}

		// Execute
		bookings, remainingTickets := purchaseTickets("Jane", "Smith", "jane.smith@example.com", 0)

		// Verify
		if remainingTickets != expectedRemainingTickets {
			t.Errorf("Expected remaining tickets: %d, got: %d", expectedRemainingTickets, remainingTickets)
		}
		if !stringSliceEqual(bookings, expectedBookings) {
			t.Errorf("Expected bookings: %v, got: %v", expectedBookings, bookings)
		}
	})
}

func purchaseTickets(firstName, lastName, email string, userTickets int) ([]string, int) {
	conferenceTickets := 50
	remainingTickets := conferenceTickets - userTickets
	bookings := []string{fmt.Sprintf("%s %s", firstName, lastName)}
	return bookings, remainingTickets
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
