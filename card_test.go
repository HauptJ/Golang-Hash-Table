package main

import (
	"testing"
)

func TestAddCard(t *testing.T) {

	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	var expectedBalance float64 = 0

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
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedBalance)
	}
}

func TestAddCardInvalid(t *testing.T) {

	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit float64 = 10
	var expectedBalance float64 = 0

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
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedBalance)
	}
}

func TestInitCard(t *testing.T) {

	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	var expectedBalance float64 = 0

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
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedBalance)
	}
}

func TestInitCardInvalid(t *testing.T) {

	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit float64 = 10
	var expectedBalance float64 = 0

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
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedBalance)
	}
}

func TestCreditCard(t *testing.T) {

	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	testCreditStr := "$10"
	var expectedNewBalance float64 = -10

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.CreditCard(testCreditStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestCreditInvalidCard(t *testing.T) {

	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit float64 = 10
	testCreditStr := "$10"
	var expectedNewBalance float64 = 0

	testCardNumStr := "1111"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.CreditCard(testCreditStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestChargeCard(t *testing.T) {
	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	testChargeStr := "$10"
	var expectedNewBalance float64 = 10

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestChargeCardInvalid(t *testing.T) {
	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit float64 = 10
	testChargeStr := "$10"
	var expectedNewBalance float64 = 0

	testCardNumStr := "1111"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestChargeCardOverLimit(t *testing.T) {
	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	testChargeStr := "$20"
	var expectedNewBalance float64 = 0

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestChargeCardNegativeAmt(t *testing.T) {
	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	testChargeStr := "-$10"
	var expectedNewBalance float64 = 0

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 == nil {
		t.Errorf("An error was not thrown when one should have been: %v", actualErr1)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}
}

func TestRetrieveBalance(t *testing.T) {
	// Arrange
	expectedStatus := "VALID"
	var expectedLimit float64 = 10
	testChargeStr := "$10"
	var expectedNewBalance float64 = 10
	expectedNewBalanceStr := "$10"

	testCardNumStr := "562246784655"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)
	actualNewBalanceStr, actualErr2 := actualCard.RetrieveBalance()

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualErr2 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr2)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}

	if actualNewBalanceStr != expectedNewBalanceStr {
		t.Errorf("Balance String was wrong, got: %s, want: %s", actualNewBalanceStr, expectedNewBalanceStr)
	}
}

func TestRetrieveBalanceInvalid(t *testing.T) {
	// Arrange
	expectedStatus := "INVALID"
	var expectedLimit float64 = 10
	testChargeStr := "$10"
	var expectedNewBalance float64 = 0
	expectedNewBalanceStr := "ERROR"

	testCardNumStr := "1111"
	testLimitStr := "$10"

	// Act
	actualCard := new(Card)
	actualErr := actualCard.InitCard(testCardNumStr, testLimitStr)
	actualErr1 := actualCard.ChargeCard(testChargeStr)
	actualNewBalanceStr, actualErr2 := actualCard.RetrieveBalance()

	// Assert
	if actualErr != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr)
	}

	if actualErr1 != nil {
		t.Errorf("An error was thrown when it should not have been: %v", actualErr1)
	}

	if actualErr2 == nil {
		t.Errorf("An error was not thrown when one should have been: %v", actualErr2)
	}

	if actualCard.status != expectedStatus {
		t.Errorf("Status was wrong, got: %s, want: %s", actualCard.status, expectedStatus)
	}

	if actualCard.limit != expectedLimit {
		t.Errorf("Limit was wrong, got: %f, want: %f", actualCard.limit, expectedLimit)
	}

	if actualCard.balance != expectedNewBalance {
		t.Errorf("Balance was wrong, got: %f, want: %f", actualCard.balance, expectedNewBalance)
	}

	if actualNewBalanceStr != expectedNewBalanceStr {
		t.Errorf("Balance String was wrong, got: %s, want: %s", actualNewBalanceStr, expectedNewBalanceStr)
	}
}
