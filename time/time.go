package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(23, time.January, 2, 2, 30, 30, 30, time.UTC)
	fmt.Println(createdTime)
}
