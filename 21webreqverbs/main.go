package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Welcome to web requests in golang")

	PerformGetRequest()
}

func PerformGetRequest() {
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