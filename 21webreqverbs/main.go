package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests in golang")

	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	fmt.Println("Doing get call")
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response content: ", string(content))

	var responseBuilder strings.Builder
	byteCount, _ := responseBuilder.Write(content)

	fmt.Println("Byte count is: ", byteCount)

	fmt.Println("Response content: ", responseBuilder.String())
}

func PerformPostJsonRequest() {
	fmt.Println("Doing post json call")
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Lets go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	fmt.Println("Doing post form call")
	const myurl = "http://localhost:8000/postform"

	// form data

	data := url.Values{}
	data.Add("firstname", "harjeevan")
	data.Add("lastname", "singh")
	data.Add("email", "singh@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
