package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Question slice
var Questions []string

// Answers Slice
var Answers []string

var currentQuestion string
var currentAnswer string

var score int

func main() {
	r := readCSV()

	for {

		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		Questions = append(Questions, record[0])
		Answers = append(Answers, record[1])

	}

	fmt.Println("Questions: ", Questions)
	fmt.Println("Answers: ", Answers)

	startGame()

	choice := ""
	fmt.Println("\nWant to Restart? ")
	fmt.Scanln(&choice)

	if choice == "Y" || choice == "y" {
		startGame()
	} else if choice == "N" || choice == "n" {
		os.Exit(0)
	} else {
		fmt.Println("\nWrong Choice. Exiting...")
		os.Exit(1)
	}

}

func startGame() {
	score = 0
	givenAnswer := ""
	for i := 0; i < len(Questions); i++ {
		currentQuestion = Questions[i]
		currentAnswer = Answers[i]
		givenAnswer = ""
		fmt.Printf("\nQuestion %d: %s \n", i+1, currentQuestion)
		fmt.Scanln(&givenAnswer)

		if givenAnswer == currentAnswer {
			score++
		}
	}
	fmt.Printf("\nYour Score: %d", score)
}

func readCSV() *csv.Reader {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	return csv.NewReader(csvFile)

}
