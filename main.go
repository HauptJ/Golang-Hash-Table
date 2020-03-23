// CCHashMap project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Println(filePath)

	scanner.Split(bufio.ScanLines)
	var streamLines []string

	for scanner.Scan() {
		streamLines = append(streamLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	processStream(streamLines)
}

func processStream(streamLines []string) {

	for _, line := range streamLines {
		fmt.Println(line)
	}

}
