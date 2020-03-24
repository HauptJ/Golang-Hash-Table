package main

import (
	"testing"
)

func TestDollarValToFloat(t *testing.T) {

	// Arrange
	var expectedFloat float64 = 10

	testString := "$10"

	// Act
	actualFloat, actualErr := DollarValToFloat(testString)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualFloat != expectedFloat {
		t.Errorf("Dollar string to int was incorrect, got: %f, want: %f", actualFloat, expectedFloat)
	}

}

func TestFloatToDollarVal(t *testing.T) {

	// Arrange
	expectedStr := "$10"

	var testFloat float64 = 10

	// Act
	actualStr, actualErr := FloatToDollarVal(testFloat)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualStr != expectedStr {
		t.Errorf("Int to dollar string was incorrect, got: %s, want: %s", actualStr, expectedStr)
	}
}
