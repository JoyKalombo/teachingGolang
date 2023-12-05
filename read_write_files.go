package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Opening a file to read
	inputFilePath := "input.txt"
	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Printing the content of the file
	fmt.Println("File Content:")
	fmt.Println(string(content))
}