package main

import "fmt"

func main()  {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages: ", languages)
	fmt.Println("JS Abbreviation is: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}