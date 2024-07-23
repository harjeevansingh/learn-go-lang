package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	fmt.Println("Days of the week are: ", days)
	fmt.Printf("Days of the week are: %v\n", days)

	for d := 0; d < len(days); d++ {
		fmt.Printf("Day %d of week is %v\n", d, days[d])
	}

	for i := range days {
		fmt.Printf("Day %d of week is %v\n", i, days[i])
	}

	for i, day := range days {
		fmt.Printf("Day %d of week is %v\n", i, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			rougueValue++
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping to LCO")
}
