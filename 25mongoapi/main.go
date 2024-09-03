package main

import (
	"fmt"
	"net/http"

	"github.com/harjeevansingh/mongoapi/router"
)

func main(){
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Run the server using 'go run main.go' command")
	http.ListenAndServe(":4000", r)
	fmt.Println("Server is running on port 4000...")
}