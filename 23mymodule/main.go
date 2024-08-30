package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	fmt.Println("Welcome to my module")
	greeter()

	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":4000", r)

	log.Fatal(http.ListenAndServe(":4000", r))
	
}

func greeter()  {
	fmt.Println("Hello there!")
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the home page")
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}