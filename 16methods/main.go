package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Structs in golang")
	// No inheritance in golang; No super or parent

	harjeevan := User{"Harjeevan", "harjeevan@go.dev", true, 16}
	fmt.Println(harjeevan)
	fmt.Printf("Harjeevan details are: %+v\n", harjeevan)
	fmt.Printf("Name is: %v\n and Email is: %v\n", harjeevan.Name, harjeevan.Email)

	harjeevan.GetStatus()
	harjeevan.NewMail()
	fmt.Printf("Name is: %v\n and Email is: %v\n", harjeevan.Name, harjeevan.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User    ) GetStatus()  {
	fmt.Println("Is user active: ", u.Status)
}

func (u User    ) NewMail()  {
	u.Email = "harjeevanNew@go.dev"
	fmt.Println("New email is: ", u.Email)
}