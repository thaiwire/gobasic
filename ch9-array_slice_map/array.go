package main

import (
	"fmt"
	_ "fmt"
)

func SampleArray1() {
	// declaration array 3 element of type string

	var fruits [3]string

	// assign value to array
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	//fruits[3] = "grape" // error: index out of range [3] with length 3

	// print loop array
	for i, fruit := range fruits {
		println("fruits[", i, "] =", fruit)
	}

	fmt.Println("element of fruits array:", fruits)
	// นับ สมาชิกใน array
	fmt.Println("Length of fruits array:", len(fruits))
	fmt.Println("Capacity of fruits array:", cap(fruits))
	fmt.Println("Fruit of index 1:", fruits[1])
	fmt.Println("Fruit of index 2:", fruits[1+1])
	index := 0
	fmt.Println("Fruit of index 0:", fruits[index])

	// array literal by loop
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("fruits[%d] = %s\n", i, fruits[i])
	}

	fmt.Println("----------------------")
	// array literal by range
	for i := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruits[i])
	}

}

func SampleArray2() {
	// ประกาศ array ขนาด 5 ช่องเประเภท string
	var carsbrands = [5]string{
		"Toyota",
		"Honda",
		"Nissan",
		"Mazda",
		"Subaru",
	}

	// print loop array
	for i, car := range carsbrands {
		println("carsbrands[", i, "] =", car)
	}

	fmt.Println("element of carsbrands array:", carsbrands)
	// นับ สมาชิกใน array
	fmt.Println("Length of carsbrands array:", len(carsbrands))
	fmt.Println("Capacity of carsbrands array:", cap(carsbrands))
	fmt.Println("Car brand of index 1:", carsbrands[1])
	fmt.Println("Car brand of index 2:", carsbrands[1+1])
	index := 0
	fmt.Println("Car brand of index 0:", carsbrands[index])

	// กรณีกำหนด array ไม่ครบทุกช่อง จะได้ค่า default ของชนิดข้อมูลนั้นๆ
	var laptops = [5]string{
		"Dell",
		"HP",
	}
	fmt.Println("element of laptops array:", laptops)

	// กรณี array ขนาด 10 ช่อง แต่กำหนดสมาชืก 5 ช่อง
	var amounts = [10]float64{
		100.50,
		200.75,
		300.25,
		400.00,
		500.10,
	}
	fmt.Println("element of amounts array:", amounts)
	// กำหนด สมาชิกใน array index ได้
	var numbers = [5]int{
		0: 10,
		2: 30,
		4: 50,
	}
	fmt.Println("element of numbers array:", numbers)

	// กำหนด array ที่มีขนาดไม่ทราบล่วงหน้า โดยใช้ ... แทนขนาดของ array
	var colors = [...]string{
		"red",
		"green",
		"blue",
		"yellow",
	}
	fmt.Println("element of colors array:", colors)
	// array append ไม่ได้ เพราะ array มีขนาดคงที่
	//colors = append(colors, "purple") // error: cannot use append on array

}
