package main

import "fmt"

func SampleSlice1() {
	var fruits []string
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Cherry")

	fmt.Println("Fruits:", fruits)

	fmt.Println("--------------")

	// เข้าถึงข้อมูลใน slice
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Second fruit:", fruits[1])
	fmt.Println("Third fruit:", fruits[2])

	fmt.Println("--------------")

	// เพิ่มข้อมูลใน slice
	fruits = append(fruits, "Grapes")
	fruits = append(fruits, "Mango")
	fmt.Println("Fruits after adding Grapes and Mango:", fruits)

	fmt.Println("--------------")

	// สร้าง slice ใหม่โดยใช้ literal
	vegetables := []string{"Carrot", "Broccoli", "Spinach"}
	fmt.Println("Vegetables:", vegetables)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	fmt.Println("Number at index 2", numbers[2])

	// add สมาชิกใน slice
	numbers = append(numbers, 6)
	fmt.Println("Numbers after adding 6:", numbers)

}
