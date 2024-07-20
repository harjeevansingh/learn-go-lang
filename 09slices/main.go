package main

import "fmt"

func main()  {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Type of fruitList is: %T", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println("Fruit list is: ", fruitList)

	fruitList = fruitList[1:3]
	fmt.Println("Fruit list after slicing: ", fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	fmt.Println("Score is: ", highScore)
} 