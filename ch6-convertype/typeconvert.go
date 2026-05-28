package main

import (
	"fmt"
	"strconv"
)

// แปลงจำนวนเต็มเป็นจำนวนทศนิยม
func TypeConvert1() {
	var i int = 100
	var f float64 = float64(i)
	fmt.Printf("i = %d, f = %.1f\n", i, f)
}

// แปลงจำนวนทศนิยมเป็นจำนวนเต็ม
func TypeConvert2() {
	var f float64 = 3.14
	var i int = int(f)
	fmt.Printf("f = %.2f, i = %d\n", f, i)
}

// แปลงจำนวนเต็มเป็นสตริง
func TypeConvert3() {
	var i int = 42
	var s string = strconv.Itoa(i) // interger to ascii
	fmt.Printf("i = %d, s = %s\n", i, s)
}

// แปลงสตริงเป็นจำนวนเต็ม
func TypeConvert4() {
	var s string = "123"
	var i int
	var err error

	i, err = strconv.Atoi(s) // ascii to interger
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return
	}
	fmt.Printf("s = %s, i = %d\n", s, i)
}

// แปลง  bool เป็น int
func TypeConvert5() {
	var b bool = false
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Printf("b = %t, i = %d\n", b, i)
}
