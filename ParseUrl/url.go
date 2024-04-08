package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://www.example.com/path/to/resource?param1=value1&param2=value2#fragment")
	if err != nil {
		panic(err)
	}

	fmt.Println("Protocol:", u.Scheme)
	fmt.Println("Hostname:", u.Hostname())
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)
}
