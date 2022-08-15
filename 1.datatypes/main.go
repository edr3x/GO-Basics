package main

import "fmt"

func main() {
	//* String
	var firstName string = "Obi-Wan"
	var lastName = "kenobi"
	var signature string

	signature = "Hello There"

	fmt.Println(firstName, lastName, "says", signature)

	name := "Obi Wan Kenobi"

	fmt.Println(name)

	//* Int

	var ageOne int = 9
	var ageTwo = 30
	ageThree := 50

	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int8 = -25
	var numTwo uint8 = 44
	numThree := int(6)

	fmt.Println(numOne, numTwo, numThree)

	//* Float

	var randomFloat1 float32 = -33.33
	var randomFloat2 float64 = 323.33890798479102748912789347298
	randomFloat3 := float32(-33.33)
	randomFloat4 := 323.33890798479102748912789347298

	fmt.Println(randomFloat1, randomFloat2, randomFloat3, randomFloat4)

	//* bool

	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	println(b3) // Prints false if not defined
	println(b1, b2, b4)
}
