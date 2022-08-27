package main

import (
	"fmt"
	"time"
)

func compute (value int){
    for i:= 0; i< value; i++{
        time.Sleep(time.Second)
        fmt.Println(i)
    }
}

func main() {
	fmt.Println("Confurrency with GoRoutines")

    // normal syncronous function call
    /*
    compute(5)
    compute(5)
    */

    // Now with goroutine
     go compute(5)
     compute(5)

    // Adding `go` keyword in front of a function make them goroutene
    // They execute at parallel i.e. we get following output

    /*
        0
        0
        1
        1
        2
        2
        3
        3
        4
        4
    */

    // now main function completes before our async function call can execute
    // As such any goroutine that is yet to complete by the end of the main function call are promptly terminated

    fmt.Scanln() // Scanln() stops the terminatin of program which thus leads to running of goroutene
}
