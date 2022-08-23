package main

import (
	"fmt"
	"sort"
)

func main() {

    /*
	yo := "Hello there!!"

	fmt.Println(strings.Contains(yo, "Hello"))
    fmt.Println(strings.ReplaceAll(yo,"ll", "uu"))
    fmt.Println(strings.ToUpper(yo))
    fmt.Println(strings.Index(yo, "re"))
    fmt.Println(strings.Split(yo," "))
    fmt.Println(yo)

    //* String library don't change the original string
   */ 

    ages:= []int{45,20,35,75,44,76,88}
    fmt.Println(ages) 
    sort.Ints(ages)
    fmt.Println(ages) // prints sorted array

    index := sort.SearchInts(ages, 75)
    fmt.Println(index)

    names := []string{"obi-wan","yoda", "anakin", "gervious", "leila"}
    sort.Strings(names)

    fmt.Println(names)
    
    fmt.Println(sort.SearchStrings(names, "yoda"))

}
