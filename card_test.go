package main

import (
	"testing"
)

func TestAddCard(t *testing.T) {

	// Arrange
	expectedStatus := "VALID"
	var expectedLimit int64 = 10
	var expectedBalance int64 = 0

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard, actualErr := AddCard(testCardNumStr, testLimitStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %d, want: %d", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %d, want: %d", actualCard.balance, expectedBalance)
	}
}

func TestAddCardInvalid(t *testing.T) {

	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit int64 = 10
	var expectedBalance int64 = 0

	testCardNumStr := "1111"
	testLimitStr := "$10"

	// Act
	actualCard, actualErr := AddCard(testCardNumStr, testLimitStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %d, want: %d", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %d, want: %d", actualCard.balance, expectedBalance)
	}
}

func TestInitCard(t *testing.T) {

	// Arrange
	expectedStatus := "VALID"
	var expectedLimit int64 = 10
	var expectedBalance int64 = 0

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %d, want: %d", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %d, want: %d", actualCard.balance, expectedBalance)
	}
}

func TestInitCardInvalid(t *testing.T) {

	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit int64 = 10
	var expectedBalance int64 = 0

	testCardNumStr := "1111"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %d, want: %d", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %d, want: %d", actualCard.balance, expectedBalance)
	}
}

func TestCreditCard(t *testing.T) {

	// Arrange

	// Act

	// Assert
}
