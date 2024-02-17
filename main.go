package main

import (
	"fmt"
	"log"
)

func CreateMessage() (string, error) {
	message := "Hello World!"
	return message, nil
}

func main() {
	message, err := CreateMessage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
