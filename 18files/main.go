package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	// Create a file
	content := "This is a file created from golang"

	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside file is: \n", string(databyte))
}


func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
	
}