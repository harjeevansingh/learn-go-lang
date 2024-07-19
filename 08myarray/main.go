package main

import "fmt"

func main()  {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	var vegList = [3 ]string{"Potato", "beans", "Tomato"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list length is: ", len(vegList))
}