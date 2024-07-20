package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	harjeevan := User{"Harjeevan", "harjeevan@mail.com", true, 23}
	fmt.Println("Harjeevan details are: ", harjeevan)
	fmt.Printf("Harjeevan details are: %+v\n", harjeevan)
	fmt.Printf("Name is %v and Email is %v\n", harjeevan.Name, harjeevan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
