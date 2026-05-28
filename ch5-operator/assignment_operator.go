package main

import "fmt"

func AssignmentOperator() {
	var x int
	x = 5
	fmt.Println(x)

	x += 3
	fmt.Println(x)

	x -= 2
	fmt.Println(x)

	x *= 2
	fmt.Println(x)

	x /= 3
	fmt.Println(x)
	
	//modulo assignment operator
	x %= 2
	fmt.Println(x)
}
