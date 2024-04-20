package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name    string `json:"coursename"`
	Class   string `json:"class"`
	RollNo  int
	Subject string
	Tags    []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json data")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", "10", 12, "maths", []string{"Hindi", "English", "Maths"}},
		{"Nodejs Bootcamp", "10", 12, "maths", []string{"Hindi", "English", "Maths"}},
		{"Nodejs Bootcamp", "10", 12, "maths", nil},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
		"coursename": "ReactJs Bootcamp",
		"class": "10",
		"RollNo": 12,
		"Subject": "maths",
		"tags": [
				"Hindi",
				"English",
				"Maths"
		]
	}`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	// some cases where you want to add data just key value pair
	var myOnlineData map[string]interface{}
	fmt.Println(len(myOnlineData))
}
