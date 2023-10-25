// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {

	s1, _ := strconv.ParseFloat(value1, 64)

	s2, _ := strconv.ParseFloat(value2, 64)

	// Calculate and return the result

	return s1 + s2
}

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println("The result is ", result)

}
