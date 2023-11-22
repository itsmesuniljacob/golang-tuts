package main

import "fmt"

func main() {
	multiSum, multiCount := addAllValues(4, 6, 7)
	fmt.Println("The total is: ", multiSum)
	fmt.Println("The total count of values: ", multiCount)
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, len(values)
}
