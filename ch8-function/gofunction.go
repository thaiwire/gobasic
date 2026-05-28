package main

import "fmt"

// simple function
func GoFunction1() {
	fmt.Println("GoFunction1")
}

// function with parameters
func GoFunction2(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("GoFunction2:", n1, n2, "Sum:", sum)
}

// function with return value
func GoFunction3(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

// multiple return values
func GoFunction4(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}
