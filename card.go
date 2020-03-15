package main

import (
	
)

type Card struct {
	status string
	limit int64
	balance int64
}


func AddCard(status string, limitStr string) (*Card, error) {
	
	c := new(Card)
	c.status =
	c.limit = 
	c.balance = 0
	return c
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
		return "ERROR", nil
	} else if (c.status == "VALID"){
		balanceStr, err :=
		if(err != nil){
			return "Failed to retrieve balance", err
		} else {
			return balanceStr, nil
		}
	} else {
		return "ERROR", errors.new("INVALID CARD STATUS")
	}
}