package main

import "fmt"

func DeferExample() {
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
}
