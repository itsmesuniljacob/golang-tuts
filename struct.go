package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Woof!!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arrf!!"
	poodle.Speak()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
