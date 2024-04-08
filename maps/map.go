package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["EN"] = "English"
	languages["HI"] = "Hindi"
	languages["FR"] = "French"
	delete(languages, "FR")
	fmt.Println(languages["EN"])

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	akash := User{"Akash", "akash.mait20@gmail.com", true, 19}
	fmt.Println(akash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
