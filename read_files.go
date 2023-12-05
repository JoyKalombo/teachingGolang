package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	inputFilePath := "input.txt"
	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File Content:")
	fmt.Println(string(content))
}