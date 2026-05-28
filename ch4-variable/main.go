package main

import "fmt"

// variable declaration
// global variable

var xmlglobal string = "outside main"

// var xml2 string = "outside main 2" // This will cause an error

func main() {

	fmt.Println(xmlglobal)
	// local variable
	var xml string = "inside main"
	fmt.Println(xml)
	// fmt.Println(xml2) // This will cause an error

	/// แบบหลายตัวแปร
	var (
		name    string = "John"
		age     int    = 30
		country string = "USA"
	)
	fmt.Println(name, age, country)

	// multiple variable declaration in one line
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	var d, e, f = 4, "hello", 3.14
	fmt.Println(d, e, f)

	//ประกาศค่าคงที่

	const pi float64 = 3.14159
	fmt.Println(pi)

	const vat = 7.0
	fmt.Println(vat)

	const name1, name2 = "Alice", "Bob"
	fmt.Println(name1, name2)

	

}
