package greeting

import "fmt"

// public function
func SayMeeting() {
	fmt.Println("Hello! from the meeting package")
}

func sayhi() {
	fmt.Println("hi! from the private meeting package")
}
