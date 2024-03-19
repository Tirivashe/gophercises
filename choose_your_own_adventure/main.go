package main

import (
	"encoding/json"
	"fmt"
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

	var story structs.Story

	err = json.NewDecoder(file).Decode(&story)
	if err != nil {
		log.Fatal("Could not parse the json file")
	}

	fmt.Println(story)
}