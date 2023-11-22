// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {

	var result float64

	v1 := convertInputToValue(input1)
	v2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(v1, v2)

	case "-":
		result = subtractValues(v1, v2)

	case "*":
		result = multiplyValues(v1, v2)

	case "/":
		result = divideValues(v1, v2)

	default:
		panic("Invalid operation")
	}

	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {

	if value2 == 0 {
		return 0
	} else {
		return value1 / value2
	}
}

func main() {
	value1 := "10"
	value2 := "0"
	operation := "/"
	result := calculate(value1, value2, operation)
	fmt.Println(result)

}
