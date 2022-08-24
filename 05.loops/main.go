package main

import "fmt"

func main() {
    x:= 0

    for x<5 { // while loop
        fmt.Println("value of x is:", x)
        x++
    }

    names := []string{"obi-wan","yoda", "anakin", "gervious", "leila"}

    for i := 0 ; i<len(names); i++{ // for loop
        fmt.Println(names[i])
    }

    for index, val := range names{
        fmt.Printf("the value at position %v is %v \n", index, val)
    }

    for _,val := range names{ // if we don't want index and only value then
        fmt.Printf("the value is %v \n", val)
    }
}
