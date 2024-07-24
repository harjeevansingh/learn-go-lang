package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://harjeevansingh.in"

func main()  {
	fmt.Println("Welcome to web requests in golang")	


	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println("Response is: ", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data from the website is: \n", string(databytes))
}