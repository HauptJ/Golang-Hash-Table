package main

import (
	"errors"

	"github.com/HauptJ/go-luhn"
)

type Card struct {
	status  string
	limit   float64
	balance float64
}

func AddCard(cardNumStr string, limitStr string) (*Card, error) {

	if cardNumStr == "" || limitStr == "" {
		return nil, errors.New("NIL status and / or limitStr passed in")
	}

	val, err := DollarValToFloat(limitStr)
	if err != nil {
		return nil, errors.New("Failed to convert limitStr to float")
	}

	c := new(Card)
	if luhn.Valid(cardNumStr) == true {
		c.status = "VALID"
	} else {
		c.status = "INVALID"
	}

	c.limit = val
	c.balance = 0
	return c, nil
}

func (c *Card) InitCard(cardNumStr string, limitStr string) error {

	if cardNumStr == "" || limitStr == "" {
		return errors.New("NIL status and / or limitStr passed in")
	}

	val, err := DollarValToFloat(limitStr)
	if err != nil {
		return errors.New("Failed to convert limitStr to float")
	}

	if luhn.Valid(cardNumStr) == true {
		c.status = "VALID"
	} else {
		c.status = "INVALID"
	}

	c.limit = val
	c.balance = 0

	return nil

}

func (c *Card) CreditCard(amountStr string) error {

	if c.status == "VALID" {
		amount, err := DollarValToFloat(amountStr)
		if err != nil {
			return errors.New("Failed to convert amountStr to float")
		}
		c.balance -= amount
	}

	return nil
}

func (c *Card) ChargeCard(amountStr string) error {

	if c.status == "VALID" {
		amount, err := DollarValToFloat(amountStr)
		if err != nil {
			return errors.New("Failed to convert amountStr to float")
		}

		if amount < 0.01 {
			return errors.New("Negative amount charged")
		}

		if c.balance+amount <= c.limit {
			c.balance += amount
		}
	}

	return nil
}

func (c *Card) RetrieveBalance() (string, error) {

	if c.status == "INVALID" {
		return "ERROR", errors.New("CARD NUMBER IS INVALID")
	} else if c.status == "VALID" {
		balanceStr, err := FloatToDollarVal(c.balance)
		if err != nil {
			return "Failed to retrieve balance", err
		} else {
			return balanceStr, nil
		}
	} else {
		return "ERROR", errors.New("NON-SUPPORTED CARD STATUS")
	}
}
