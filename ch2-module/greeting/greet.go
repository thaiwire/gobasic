package greeting

import (
	"fmt"

	personal "github.com/thaiwire/samplego/greeting/internal"
)

// public function
func SayGreetingHello() {
	// call private function from the same package
	personal.PrintPersonal()

	fmt.Println("Hello! from the greeting package")
	// private function can be called from within the package
	sayhi()
}
