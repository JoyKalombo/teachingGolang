package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	myFile := "output.txt"

	content, err := ioutil.ReadFile(myFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	oldContent := "I am going to write this to the file"
	newContent := "I am going to replace the output file with this"
	modifiedContent := []byte(strings.ReplaceAll(string(content), oldContent, newContent))

	err = ioutil.WriteFile(myFile, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("The file content was successfully modified :-)")
}
