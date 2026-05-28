package main

import (
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/thaiwire/samplego/greeting"
)

func main() {
	fmt.Println("Welcome to Go!")
	fmt.Printf("The square root of 2 is %g\n", math.Sqrt(2))

	// generte a new UUID
	var id string = uuid.New().String()

	fmt.Printf("A new UUID: %s\n", id)

	greeting.SayGreetingHello()
	greeting.SayMeeting()
	// private function cannot be called from outside the package
	//greeting.sayhi() // this will cause an error
}
