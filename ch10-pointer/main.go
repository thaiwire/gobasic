package main

import "fmt"

func main() {
	//SamplePointer1()
	age := 30
	println("Age before change:", age)
	changeAge(&age) // ส่ง pointer ของ age ไปยังฟังก์ชัน changeAge
	println("Age after change:", age)

	fmt.Println("------------------	")

	SamplePointer2()
}
