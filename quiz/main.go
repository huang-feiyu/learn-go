package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func startQuiz(expressions [][]string) (corrcet, all int) {
	all = len(expressions)
	corrcet = 0

	for i, expression := range expressions {
		fmt.Printf("Problem #%d: %s = ", i+1, expression[0])
		var answer string
		if _, err := fmt.Scan(&answer); err != nil {
			log.Fatal("Unable to read answer", err)
		}
		// remove spaces from answer
		if strings.ReplaceAll(answer, " ", "") == expression[1] {
			corrcet++
		}
	}

	return
}

func main() {
	file := flag.String("csv", "problems.csv", "a csv file in the format of \"question,answer\"")
	flag.Parse()
	expressions := readCsvFile(*file)

	correct, all := startQuiz(expressions)
	fmt.Println("You scored", correct, "out of", all, "correct")
}
