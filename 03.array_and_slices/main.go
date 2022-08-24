package main

import "fmt"

func main() {

	//	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var arr = [5]int{1, 2, 3, 4, 5}

	names := [4]string{"jhon", "doe", "obi-wan", "kenobi"}

	fmt.Println(arr, len(arr))
	fmt.Println(names, len(names))

	//* Slices(use arrays under the hood)

	var scores = []int{445, 232, 435}
	scores[0] = 670
	fmt.Println(scores, len(scores))
	scores = append(scores, 85) //* can't do append with array but can do with slices

	fmt.Println(scores, len(scores))

	//slice ranges

	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]

	fmt.Println(range1, "\n", range2, "\n", range3)

	range1 = append(range1, "broda")

	fmt.Println(range1)

}
