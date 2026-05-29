package main

import _ "fmt"

func main() {
	// สร้าง channel ที่รับค่าเป็น string
	ch := make(chan string)
	// สร้าง goroutine ที่ส่งข้อความไปยัง channel
	go func() {
		ch <- "Hello, Channel!" // ส่งข้อความ "Hello, Channel!" ไปยัง channel
	}()
	// รับข้อความจาก channel และพิมพ์ออกมา
	msg := <-ch
	println(msg)

}
