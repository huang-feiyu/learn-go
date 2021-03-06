package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readCsvFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	// ensure to close the file
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	// csv line: question,answer
	return records
}

func startQuiz(expressions [][]string, timer *time.Timer) (correct, all int) {
	correct, all = 0, 0

	for i, expression := range expressions {
		fmt.Printf("Problem #%d: %s = ", i+1, expression[0])
		answerChannel := make(chan string)
		all++
		go func() {
			var answer string
			if _, err := fmt.Scan(&answer); err != nil {
				log.Fatal("Unable to read answer", err)
			}
			answerChannel <- strings.TrimSpace(answer)
		}()
		select {
		case <-timer.C:
			return
		case answer := <-answerChannel:
			// remove spaces from answer
			if answer == expression[1] {
				correct++
			}
		}
	}

	return
}

func main() {
	file := flag.String("csv", "problems.csv", "a csv file in the format of \"question,answer\"")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	expressions := readCsvFile(*file)

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct, all := startQuiz(expressions, timer)

	fmt.Println("\nYou scored", correct, "out of", all)
}
