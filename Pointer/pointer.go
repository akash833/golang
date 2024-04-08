package main

import "fmt"

func main() {
	hello := 23
	ptr := &hello
	fmt.Println(ptr, " ", *ptr)
}
