package main

import "fmt"

// Write your answer here, and then test your code.

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here
	elementValue := 100 / float64(len(items))
	for _, fruit := range items {
		result[fruit] = elementValue
	}
	return result
}

func main() {
	slice := []string{"apple", "banana"}
	result := convertToMap(slice)
	fmt.Println("The result is ", result)
}
