package main

import (
	"fmt"
	_ "fmt"
)

func SamplePointer1() {
	// บ้านเลขที่ 10
	value := 10

	// ตัวแปร pointer ที่เก็บที่อยู่ของ value
	var pointer *int = &value
	// แสดง Address ของ pointer
	println("Address of pointer:", pointer) // แสดงที่อยู่ของ pointer
	// แสดงค่าของ value ผ่าน pointer
	println("Value:", *pointer) // ใช้ * เพื่อ dereference pointer

	// เปลี่ยนค่าของ value ผ่าน pointer
	*pointer = 20
	println("Updated Value:", value)                    // แสดงค่าที่อัพเดตของ value
	println("Updated Value through pointer:", *pointer) // แสดงค่าที่อัพเดตผ่าน pointer

	fmt.Println("------------------	")
	// pointer กับ slice
	slice := []int{1, 2, 3}
	pointerSlice := &slice
	fmt.Println("Original slice:", slice)
	fmt.Println("Pointer to slice:", pointerSlice)
	fmt.Println("Value through pointer:", *pointerSlice)

	// เปลี่ยนค่าของ slice ผ่าน pointer
	(*pointerSlice)[0] = 10
	fmt.Println("Updated slice:", slice)
	fmt.Println("Updated value through pointer:", *pointerSlice)
}

func changeAge(age *int) { // รับ pointer ที่ชี้ไปยังตัวแปร age
	*age += 1
}

// ฟังก์ชันตัวอย่างที่สองสำหรับการใช้งาน pointer กับ struct

type Room struct {
	RoomNumber int
	BedColor   string
}

// ฟังชัน เปลื่ยนสีของเตียงในห้อง

func changeBedColor(keyToRoom *Room, newColor string) {
	keyToRoom.BedColor = newColor // ใช้ dot notation เพื่อเปลี่ยนสีของเตียงผ่าน pointer
}

func SamplePointer2() {
	// สร้าง instance ของ Room
	room101 := Room{
		RoomNumber: 101,
		BedColor:   "Yellow",
	}
	fmt.Printf("Before change: Room %d has bed color %s\n", room101.RoomNumber, room101.BedColor)
	changeBedColor(&room101, "Blue") // ส่ง pointer ของ room101 ไปยังฟังก์ชัน changeBedColor พร้อมกับสีใหม่
	fmt.Printf("After change: Room %d has bed color %s\n", room101.RoomNumber, room101.BedColor)
}
