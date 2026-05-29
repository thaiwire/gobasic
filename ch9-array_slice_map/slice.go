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
	numbers = append(numbers, 7)
	fmt.Println("Numbers after adding 6 and 7:", numbers)

	// ตัด slice โดยใช้ slice operation
	subSlice1 := numbers[1:4] // ตัด slice ตั้งแต่ index 1 ถึง 3 (ไม่รวม index 4)
	subSlice2 := numbers[:3]  // ตัด slice ตั้งแต่ index 0 ถึง 2 (ไม่รวม index 3)
	subSlice3 := numbers[2:]  // ตัด slice ตั้งแต่ index 2 ถึงสุดท้าย
	fmt.Println("Sub-slice of numbers (index 1 to 3):", subSlice1)
	// แสดง length
	fmt.Println("Length of sub-slice 1:", len(subSlice1))
	// แสดง capacity
	fmt.Println("Capacity of sub-slice 1:", cap(subSlice1))

	fmt.Println("Sub-slice of numbers (index 0 to 2):", subSlice2)
	fmt.Println("Sub-slice of numbers (index 2 to end):", subSlice3)

}
