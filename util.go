package main

import (
	"errors"
	"strconv"
	"strings"
)

func DollarValToInt(amountStr string) (int64, error) {
	if amountStr == "" {
		return 0, errors.New("Empty amountStr passed in")
	}

	val, err := strconv.ParseInt(strings.Replace(amountStr, "$", "", -1), 10, 64)
	if err != nil {
		return 0, errors.New("Failed to convert string to int")
	}

	return val, nil
}

func IntToDollarVal(amount int64) (string, error) {

	return "$" + strconv.FormatInt(amount, 10), nil
}
