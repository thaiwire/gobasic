package main

import "fmt"

func LogicalOperator() {
	var x bool
	x = true
	fmt.Println(x && true)
	fmt.Println(x || false)
	fmt.Println(!x)

	var y bool
	y = !true
	fmt.Println(y)
	
	// or
	y = (true || false)
	fmt.Println(y)
}
