package main

import "fmt"

const LoginToken string = "hjkhkjhk"

func main()  {
	var username string = "harjeevan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Small Variable is of type: %T \n", smallVal)
	
	var smallFloat float64 = 255.32434343
	fmt.Println(smallFloat)
	fmt.Printf("Small float is of type: %T \n", smallFloat)

	// Default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Another Variable is of type: %T \n", anotherVariable)

	// Implicit type
	var website = "harjeevansingh.in"
	fmt.Println(website)

	// No var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type: %T \n", LoginToken)
}