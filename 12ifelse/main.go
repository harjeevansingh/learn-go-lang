package main

import "fmt"

func main()  {
	fmt.Println("Welcome to if else in golang")

	loginCount := 10

	if loginCount < 10 {
		fmt.Println("Regular User")
	} else if loginCount > 10 {
		fmt.Println("Watch out! You are a star")
	} else {
		fmt.Println("Call security!")
	}

	if loginCount := 10; loginCount < 10 {
		fmt.Println("Regular User")
	} else {
		fmt.Println("Call security!")
	}
	
}