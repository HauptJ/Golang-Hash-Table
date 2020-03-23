package main

import (
	"errors"
	"strconv"
	"strings"
)

func DollarValToFloat(amountStr string) (float64, error) {
	if amountStr == "" {
		return 0, errors.New("Empty amountStr passed in")
	}

	val, err := strconv.ParseFloat(strings.Replace(amountStr, "$", "", -1), 10)
	if err != nil {
		return 0, errors.New("Failed to convert string to float")
	}

	return val, nil
}

func FloatToDollarVal(amount float64) (string, error) {

	return "$" + strconv.FormatFloat(amount, 'f', -1, 64), nil
}
