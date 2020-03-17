package main

import (
	"errors"

	"github.com/HauptJ/go-luhn"
)

type Card struct {
	status  string
	limit   int64
	balance int64
}

func AddCard(cardNumStr string, limitStr string) (*Card, error) {

	if cardNumStr == "" || limitStr == "" {
		return nil, errors.New("NIL status and / or limitStr passed in")
	}

	val, err := DollarValToInt(limitStr)
	if err != nil {
		return nil, errors.New("Failed to convert limitStr to int")
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

	val, err := DollarValToInt(limitStr)
	if err != nil {
		return errors.New("Failed to convert limitStr to int")
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
		amount, err := DollarValToInt(amountStr)
		if err != nil {
			return errors.New("Failed to convert amountStr to int")
		}
		c.balance -= amount
	}

	return nil
}

func (c *Card) ChargeCard(amountStr string) error {

	if c.status == "VALID" {
		amount, err := DollarValToInt(amountStr)
		if err != nil {
			return errors.New("Failed to convert amountStr to int")
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
		balanceStr, err := IntToDollarVal(c.balance)
		if err != nil {
			return "Failed to retrieve balance", err
		} else {
			return balanceStr, nil
		}
	} else {
		return "ERROR", errors.New("NON-SUPPORTED CARD STATUS")
	}
}
