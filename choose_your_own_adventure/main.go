package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Tirivashe/choose_your_own_adventure/structs"
)

func main() {

	file, err := os.Open("adventure.json")
	if err != nil {
		log.Println("An error occured opening the file:", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing the file:", err)
			os.Exit(3)
		}
	}()

	bookBytes, err :=  io.ReadAll(file)
	if err != nil {
		log.Println("Error reading from the file: ", err)
		os.Exit(2)
	}

	var book structs.Chapter
	err = json.Unmarshal(bookBytes, &book)
	if err != nil {
		log.Fatal("Could not parse the json data inside the book")
	}

	fmt.Println(book)
}