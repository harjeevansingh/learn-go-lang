package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCOurses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js", "react"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js", "react", "node"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abc123", nil},
	}

	// Package this data in JSON
	finalJson, err := json.MarshalIndent(lcoCOurses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println("Json data is: ", string(finalJson))
}

func DecodeJson()  {
	fmt.Println("Decoding json")
	jsonDataFromWeb := []byte(`
	 {
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Password": "abc123",
		"Tags": [
			"web-dev",
			"js",
			"react"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("Course is: %#v\n", lcoCourse)
	} else {
		fmt.Println("Json is invalid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("Course is: %#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is: %v \t and value is: %v \t and type is: %T\n", key, value, value)
	}


}
