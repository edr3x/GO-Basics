package main

import "fmt"

func main() {
	name := "Yoda"
	race := "Jedi"
	age := 200

	//* Printf (Formatted strings)
	fmt.Printf("Name of mine is %v and im a %v \n", name, race) //* Default value
	fmt.Printf("Name is %q and age is %d \n", name, age)        //* quoted value with decimal

	//* Sprintf (save formatted strings)

	savedStr := fmt.Sprintf("Name is %q and age is %d \n", name, age)
	fmt.Println("the saved string is:", savedStr)

}
