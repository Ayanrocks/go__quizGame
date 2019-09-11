package main

import (
	"bufio"
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

}

func startGame() {
	score = 0
	for i := 0; i < len(Questions); i++ {
		currentQuestion = Questions[i]
		currentAnswer = Answers[i]

		fmt.Printf("\nQuestion %d: %s \n", i, currentQuestion)
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		fmt.Printf("%T, %T", answer, currentAnswer)
		if answer == currentAnswer {
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
