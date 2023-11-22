package main

import "fmt"

func main() {

	sum := add(1, 2)
	multiSum, multiCount := addAllValues(4, 6, 7)
	fmt.Println("The total is: ", multiSum)
	fmt.Println("The total count of values: ", multiCount)
}

func add(value1, value2 int) int {

}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, len(values)
}
