// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {

	s1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println("Value must be a number")
	}

	s2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println("Value must be a number")
	}

	// Calculate and return the result

	return s1 + s2
}

func main() {
	value1 := "xyz"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println("The result is ", result)

}
