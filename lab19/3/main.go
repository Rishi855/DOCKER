package main

import (
	"fmt"
)

func main() {
	// Create a map to store the number of occurrences of each character
	charCount := make(map[rune]int)

	// Define the input string
	input := "hello world"

	// Iterate over each character in the string
	for _, char := range input {
		// Increment the count for the character
		charCount[char]++
	}

	// Print the character counts
	fmt.Println("Character counts in 'hello world':")
	for char, count := range charCount {
		fmt.Printf("'%c': %d\n", char, count)
	}
}
