package main

import "fmt"


func CompareOperator() {
	var x bool
	x = true
	fmt.Println(x == true)
	fmt.Println(x != true)
	// fmt.Println(x > true) // Invalid operation: cannot compare bool with bool using '>'
	// fmt.Println(x < true) // Invalid operation: cannot compare bool with bool using '<'
	// fmt.Println(x >= true) // Invalid operation: cannot compare bool with bool using '>='
	// fmt.Println(x <= true) // Invalid operation: cannot compare bool with bool using '<='
}
