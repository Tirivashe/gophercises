package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csvfile", "problems.csv", "Takes the CSV file and returns it")
	timeLimit := flag.Int("limit", 30, "Sets the time limit to finish the quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalf("Could not read the provided file: %v", err)
	}
	
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Could not read records: %v", err)
	}

	correctAnswers := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	// go func() {
	// 	time.Sleep(time.Duration(limit) * time.Second)
	// 	fmt.Printf("\nTime's up! You got %d out of %d\n", correctAnswers, len(records))
	// 	os.Exit(0)
	// }()

	for _, record := range records {
		question, answer := record[0], record[1]
		fmt.Printf("%v = ", question)
		input := make(chan string)
		go func(){
			var userAnswer string
			_, err := fmt.Scanln(&userAnswer)
			if err != nil {
				log.Fatalf("Could not read user input: %v", err)
				os.Exit(1)
			}
			input <- userAnswer
		}()
		select {
			case <-timer.C:
				fmt.Printf("\nTime's up! You got %d out of %d\n", correctAnswers, len(records))
				return
			case userInput:= <- input:
				if userInput == answer {
					correctAnswers ++
				}
		}
	}
	fmt.Printf("Finish! You got %d out of %d\n", correctAnswers, len(records))
}