package main

import "fmt"

func main() {
	// akash := User{Name: "Akash Gupta", Email: "akash.mait20@gmail.com", Age: 23}
	// fmt.Println(akash.getName())

	// defer
	fmt.Println("Hello My name is Akash gupta")
	defer fmt.Println("Hllloooo")
	fmt.Println("Aur kya haal ha")
}

type User struct {
	Name  string
	Email string
	Age   int // Age should be of type int, not string
}

func (u User) getName() string {
	return u.Name
}
