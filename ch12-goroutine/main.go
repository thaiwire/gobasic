package main

import (
	"fmt"
	_ "fmt"
	"time"
)

func brewCoffee() {
	fmt.Println("Start brewCoffee")
	time.Sleep(2 * time.Second) // ใช้เวล่า 2 วินาที
	fmt.Println("Finish brewCoffee")
}

func bakeBread() {
	fmt.Println("Start bakeBread")
	time.Sleep(3 * time.Second) // ใช้เวล่า 3 วินาที
	fmt.Println("Finish bakeBread")
}

func washDishes() {
	fmt.Println("Start washDishes")
	time.Sleep(1 * time.Second) // ใช้เวล่า 1 วินาที
	fmt.Println("Finish washDishes")
}

func main() {
	go brewCoffee() // เริ่มฟังก์ชัน brewCoffee ใน goroutine
	go bakeBread()  // เริ่มฟังก์ชัน bakeBread ใน goroutine
	go washDishes() // เริ่มฟังก์ชัน washDishes ใน goroutine

	// รอให้ goroutine ทั้งหมดทำงานเสร็จ
	time.Sleep(4 * time.Second)
	fmt.Println("All tasks completed")
}
