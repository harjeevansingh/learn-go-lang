package main

import (
	"fmt"
	"net/url"
)

// const myUrl string = "https://harjeevansingh.in:3000"
const myUrl string = "https://lco.dev:3000/learn?course=golang&paymentid=123"

func main()  {
	fmt.Println("Welcome to urls in golang")
	fmt.Println("URL is: ", myUrl)


	// Parsing the url
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme is: ", result.Scheme) // https
	fmt.Println("Host is: ", result.Host) 
	fmt.Println("Path is: ", result.Path) 
	fmt.Println("Port is: ", result.Port()) // 3000
	fmt.Println("RawQuery is: ", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of query params is: %T\n", queryParams)

	fmt.Println("Query params are: ", queryParams)
	fmt.Println("Payment ID is: ", queryParams["paymentid"])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=harjeevan",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another URL is: ", anotherUrl)

}