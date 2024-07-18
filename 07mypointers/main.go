package main

import "fmt"

func main()  {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr) 

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pointer is: ", ptr)
	fmt.Println("Value of inside the pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of myNumber is: ", myNumber)
}