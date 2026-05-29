package main

import "fmt"

func SampleMap1() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	println(m["a"])
	println(m["b"])
	println(m["c"])

	scores := make(map[string]int)
	scores["Alice"] = 90
	scores["Bob"] = 85
	scores["Charlie"] = 92

	println("Alice's score:", scores["Alice"])
	println("Bob's score:", scores["Bob"])
	println("Charlie's score:", scores["Charlie"])

	// ตรวจสอบว่ามี key อยู่ใน map หรือไม่
	if score, ok := scores["Alice"]; ok {
		println("Alice's score is", score)
	} else {
		println("Alice's score not found")
	}

	fmt.Println("Score of Alice", scores["Alice"])

	// ตรวจสอบว่ามี key อยู่ใน map หรือไม่
	if score, exists := scores["Bob"]; exists {
		fmt.Println("Bob's score is", score)
	} else {
		fmt.Println("Bob's score not found")
	}

	// ลบ key "Charlie" ออกจาก map
	delete(scores, "Charlie")
	// ตรวจสอบว่ามี key "Charlie" อยู่ใน map หรือไม่
	if score, exists := scores["Charlie"]; exists {
		fmt.Println("Charlie's score is", score)
	} else {
		fmt.Println("Charlie's score not found")
	}

	// ลูปผ่าน map และพิมพ์ key และ value
	for name1, score := range scores {
		fmt.Printf("%s's score is %d\n", name1, score)
	}

	// สร้าง map ที่มีค่าเริ่มต้น
	defaultScores := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
	}
	fmt.Println("Default Scores:", defaultScores)

	grades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "A",
	}
	fmt.Println("Grades:", grades)

}
