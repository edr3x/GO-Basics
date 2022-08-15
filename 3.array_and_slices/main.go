package main

import "fmt"

func main() {

	//	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var arr = [...]int{1, 2, 3, 4, 5}

	names := [2]string{"jhon", "doe"}
	names[1] = "dont"

	fmt.Println(arr, len(arr))
	fmt.Println(names, len(names))

	//* Slices

	// var scores = []int{445, 232, 435}
}
