package main

import "fmt"

func ForLoop1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
}

// for like While loop
func ForLoop2() {
	i := 0
	for i < 5 {
		fmt.Println("Iteration:", i)
		i++
	}
}

// Infinite loop
func ForLoop3() {

	i := 1
	for {
		fmt.Println("This will run forever!", i)
		// if i >= 5000 {
		// 	break
		// }

		if i == 1000 {
			fmt.Println("Reached 1000, skipping to 2000")
			i = 2000
			continue
		}
		i++

	}
}
