// CCHashMap project main.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := "NONE"
	scanner := bufio.NewScanner(os.Stdin)
	if len(os.Args) == 2 {
		filePath = os.Args[1]

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	scanner.Split(bufio.ScanLines)
	var streamLines []string

	for scanner.Scan() {
		streamLines = append(streamLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	res, err := processStream(streamLines)

	if err != nil {
		log.Println(err)
	}

	for _, line := range res {
		fmt.Println(line)
	}
}

func processStream(streamLines []string) ([]string, error) {

	cards := make(map[string]*Card)
	var res []string

	for _, line := range streamLines {
		fields := strings.Fields(line)
		if fields[0] == "Add" {
			card := new(Card)
			cards[fields[1]] = card
			cards[fields[1]].InitCard(fields[2], fields[3])
		} else if fields[0] == "Charge" {
			cards[fields[1]].ChargeCard(fields[2])
		} else if fields[0] == "Credit" {
			cards[fields[1]].CreditCard(fields[2])
		} else {
			return nil, errors.New("Invalid option")
		}
	}

	for name, _ := range cards {
		balance, _ := cards[name].RetrieveBalance()
		res = append(res, fmt.Sprintf("%s: %s", name, balance))
	}

	return res, nil
}
