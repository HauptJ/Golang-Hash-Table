package main

import (
	"testing"
)

func TestDollarValToInt(t *testing.T) {

	// Arrange
	var expectedInt int64 = 10

	testString := "$10"

	// Act
	actualInt, actualErr := DollarValToInt(testString)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualInt != expectedInt {
		t.Errorf("Dollar string to int was incorrect, got: %d, want: %d", actualInt, expectedInt)
	}

}

func TestIntToDollarVal(t *testing.T) {

	// Arrange
	expectedStr := "$10"

	var testInt int64 = 10

	// Act
	actualStr, actualErr := IntToDollarVal(testInt)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualStr != expectedStr {
		t.Errorf("Int to dollar string was incorrect, got: %s, want: %s", actualStr, expectedStr)
	}
}
