package main

import "fmt"

func main() {
	fruits := []string{"Mango", "Apple", "Banana", "Graphs"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for i := range fruits {
		fmt.Println(fruits[i])
	}

	for index, day := range fruits {
		fmt.Printf("The %v is present at %v\n", day, index)
	}
	fmt.Println(proAdder(23, 45, 65, 67, 87))
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
