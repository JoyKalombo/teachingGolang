package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// countWords counts the frequency of each word in a paragraph
func countWords(paragraph string) map[string]int {
	wordFrequency := make(map[string]int)
	wordStart := -1

	for i, char := range paragraph {
		if isWhitespace(char) || isPunctuation(char) {
			if wordStart != -1 {
				word := paragraph[wordStart:i]
				wordFrequency[word]++
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
	}

	// Check for the last word in the paragraph
	if wordStart != -1 {
		word := paragraph[wordStart:]
		wordFrequency[word]++
	}

	return wordFrequency
}

// isWhitespace checks if a character is a whitespace character
func isWhitespace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

// isPunctuation checks if a character is a punctuation character
func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ';'
}

// replaceWord replaces a specific word in the paragraph with another word
func replaceWord(paragraph, oldWord, newWord string) string {
	// Use strings.Replace to replace all occurrences of oldWord with newWord
	return strings.Replace(paragraph, oldWord, newWord, -1)
}

func main() {
	// Example usage

	// Opening a file to read
	inputFilePath := "paragraph.txt"
	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	paragraph := string(content)

	wordFrequency := countWords(paragraph)

	// Print the word frequency map
	fmt.Println("Word Frequency:")
	for word, frequency := range wordFrequency {
		fmt.Printf("%s: %d\n", word, frequency)
	}

	// Allow the user to input a word to find its frequency
	var userInput string
	fmt.Print("\nEnter a word to find its frequency: ")
	fmt.Scan(&userInput)

	// Print the frequency of the user input word
	if frequency, found := wordFrequency[userInput]; found {
		fmt.Printf("The word '%s' appears %d times in the paragraph.\n", userInput, frequency)
	} else {
		fmt.Printf("The word '%s' does not appear in the paragraph.\n", userInput)
	}

	// Allow the user to input old and new words for replacement
	var oldWord, newWord string
	fmt.Print("\nEnter the word to replace: ")
	fmt.Scan(&oldWord)
	fmt.Print("Enter the new word: ")
	fmt.Scan(&newWord)

	// Replace the specified word in the paragraph
	updatedParagraph := replaceWord(paragraph, oldWord, newWord)

	// I will write new things to the file

	outputFilePath := "paragraph.txt"
	err = ioutil.WriteFile(outputFilePath, []byte(updatedParagraph), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Print the updated paragraph
	fmt.Println("\nUpdated Paragraph:")
	fmt.Println(updatedParagraph)
}

// Can we do it so that a word chosen is then changed?

// To run a front end version, use the following packages:
// "http" and "templates"