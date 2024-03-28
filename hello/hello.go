package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("Gladys")
	if err {
		log.Fatal(err)
	}
	fmt.Println(message)
}
