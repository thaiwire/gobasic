package main

import "fmt"

func IfCondition() {
	var a int = 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else if a == 5 {
		fmt.Println("a is equal to 5")
	} else {
		fmt.Println("a is less than 5")
	}
}

func IfElseIfCondition() {
	var score int = 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

func SwitchCondition() {
	var day int = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

func SwithMultipleCase() {
	var month int = 4
	switch month {
	case 1, 2, 3:
		fmt.Println("Spring")
	case 4, 5, 6:
		fmt.Println("Summer")
	case 7, 8, 9:
		fmt.Println("Autumn")
	case 10, 11, 12:
		fmt.Println("Winter")
	default:
		fmt.Println("Invalid month")
	}
}
