package main

import "fmt"

func main()  {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()
	fmt.Println("Result is : ", adder(2, 3))

	total, _ := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Total is: ", total)

}

func greeter() {
	fmt.Println("Hello there!")
}

func greeterTwo() {
	fmt.Println("Hello there again!")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hello there"
}