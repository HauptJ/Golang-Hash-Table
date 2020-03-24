package main

import (
	"testing"
)

func TestProcessStream(t *testing.T) {
	// Arrange
	testStream := []string{"Add Tom 4111111111111111 $1000",
		"Add Lisa 5454545454545454 $3000", "Add Quincy 1234567890123456 $2000",
		"Charge Tom $500", "Charge Tom $800", "Charge Lisa $7",
		"Credit Lisa $100", "Credit Quincy $200"}

	// Act
	_, actualErr := processStream(testStream)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}
}
