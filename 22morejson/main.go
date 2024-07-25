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

	EncodeJson()
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
