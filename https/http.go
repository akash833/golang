package main

import (
	"fmt"
	"net/http"
)

const url = "http://localhost:8000"

func main() {
	fmt.Println("Welcome to http")

	// simple get request
	content, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
