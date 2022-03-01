package main

import (
	"fmt"

	"github.com/intelliviz.com/basics/src/simple/numbers"
)

// Test
func main() {
	// var age int = 50	// full declaration
	var age = 50 // don't need type-go can infer type from assigned value

	var num int = numbers.GetNumberFive()
	fmt.Printf("Hello World: %d %d\n", num, age)

	// short declaration operator-short hand method for variable declaration
	// only works in block scope
	s := "Learning Go-Lang"
	fmt.Println(s)

	// blank variable-declare a variable without using it
	_ = "a value"

	car, cost := "Toyota", 25000
	fmt.Println(car, cost)

	var a = 4
	var b = 5.2
	// a = b can't do this
	a = int(b)
	fmt.Println(a, b)
}
