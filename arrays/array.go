package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "graphs", "pineapple")
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore = append(highScore, 555, 666, 777)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slice based on index

	var courses = []string{"reactjs", "nodejs", "javascript", "java"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
