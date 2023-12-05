package main
import (
	"fmt"
	"io/ioutil"
)

func main() {

	data := []byte("I am going to write this to the file")

	outputFilePath := "output.txt"
	err := ioutil.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("The data has successfully been written to the file.")
}