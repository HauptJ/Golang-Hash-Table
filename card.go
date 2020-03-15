package main

import (
	"github.com/HauptJ/go-luhn"
)

type Card struct {
	status string
	limit int64
	balance int64
}


func AddCard(cardNumStr string, limitStr string) (*Card, error) {
	
	if(status == nil || limitStr == nil){
		return nil, errors.new("NIL status and / or limitStr passed in")
	}
	
	val, err := dollarValToInt(limitStr)
	if(err != nil){
		return nil, errors.new("Failed to convert limitStr to int")
	}
	
	c := new(Card)
	if(luhn.Valid(cardNumStr) == true){
		c.status = "VALID"
	} else {
		c.status = "INVALID"
	}
	
	c.limit = val
	c.balance = 0
	return c
}


func (c *Card) InitCard(cardNumStr string, limitStr string) error {
	
	if(status == nil || limitStr == nil){
		return nil, errors.new("NIL status and / or limitStr passed in")
	}
	
	val, err := dollarValToInt(limitStr)
	if(err != nil){
		return nil, errors.new("Failed to convert limitStr to int")
	}
	
	if(luhn.Valid(cardNumStr) == true){
		c.status = "VALID"
	} else {
		c.status = "INVALID"
	}
	
	c.limit = val
	c.balance = 0
	
}


func (c *Card) CreditCard(string amountStr) error {
	
	if(c.status == "VALID"){
		amount := 
		c.balance -= amount
	}
}


func (c *Card) ChargeCard(string amountStr) error {
	
	if(c.status == "VALID"){
		amount := 
		if(c.balance + amount <= c.limit){
			c.balance += amount
		}
	}
}


func (c *Card) RetrieveBalance() (string, error) {
	
	if(c.status == "INVALID"){
		return "ERROR", errors.new("CARD NUMBER IS INVALID")
	} else if (c.status == "VALID"){
		balanceStr, err :=
		if(err != nil){
			return "Failed to retrieve balance", err
		} else {
			return balanceStr, nil
		}
	} else {
		return "ERROR", errors.new("NON-SUPPORTED CARD STATUS")
	}
}