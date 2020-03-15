package main

func dollarValToInt(amountStr string) (int64, error) {
	if amountStr == nil {
		return nil, errors.new("Empty amountStr passed in")
	}

	val, err := ParseInt(strings.Replace(amountStr, "$", "", -1))
	if err != nil {
		return nil, errors.new("Failed to convert string to int")
	}

	return val, nil
}

func IntToDollarVal(amount int64) (string, error) {
	if ammountStr == nil {
		return nil, errors.new("Empty amount passed in")
	}

	return "$" + FormatInt(amount, 10), nil
}
