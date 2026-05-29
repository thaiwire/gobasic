package main

import "fmt"

//กำหนด struct ชื่อ Person ที่มีฟิลด์ Name และ Age
type Person struct {
	Name string
	Age  int
}

// ฟังชัน sayHello เป็นเมธอดของ struct Person ที่พิมพ์ข้อความทักทาย
func (p Person) sayHello() {
	println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

// ฟังชั่น  HaveBirthday เป็นเมธอดของ struct Person ที่เพิ่มอายุของบุคคลขึ้น 1 ปี
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	//สร้างตัวแปร p1 ของประเภท Person และกำหนดค่าให้กับฟิลด์
	p1 := Person{Name: "Alice", Age: 30}
	//สร้างตัวแปร p2 ของประเภท Person และกำหนดค่าให้กับฟิลด์
	p2 := Person{Name: "Bob", Age: 25}
	//แสดงข้อมูลของ p1 และ p2
	println("Person 1:", p1.Name, "Age:", p1.Age)
	println("Person 2:", p2.Name, "Age:", p2.Age)

	fmt.Println("------------------") // พิมพ์บรรทัดว่างเพื่อแยกส่วนของข้อมูลและผลลัพธ์ของเมธอด
	// เรียกใช้เมธอด sayHello ของ p1 และ p2
	p1.sayHello()
	p2.sayHello()

	fmt.Println("------------------") // พิมพ์บรรทัดว่างเพื่อแยกส่วนของข้อมูลและผลลัพธ์ของเมธอด
	// เรียกใช้เมธอด HaveBirthday ของ p1 และ p2
	p1.HaveBirthday()
	p2.HaveBirthday()

	// แสดงข้อมูลของ p1 และ p2 หลังจากมีวันเกิด
	println("Person 1 after birthday:", p1.Name, "Age:", p1.Age)
	println("Person 2 after birthday:", p2.Name, "Age:", p2.Age)
}
