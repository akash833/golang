package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This will write in the file"
	file, err := os.Create("./hello.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()

	readFile, err := os.ReadFile("./hello.txt")
	fmt.Println(string(readFile))
}
