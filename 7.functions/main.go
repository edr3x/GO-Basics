package main

import (
	"fmt"
	"math"
	"strings"
)

func greeting(n string){
    fmt.Printf("Good morning %v \n",n)
}

func bye(n string){
    fmt.Printf("Goodbye %v\n",n)
}

func cycleNames(n []string, f func(string)){
    for _, val := range n{
        f(val)
    }
}

func circleArea(r float64) float64{
    return math.Pi * r * r
}

// Multiple Return type functions

func getInitials(n string) (string, string){
    s := strings.ToUpper(n)
    names := strings.Split(s," ")
    
    var initials []string

    for _,val := range names{
        initials = append(initials, val[:1]) // gets first letter of a string
    }
    
    if len(initials) > 1{
        return initials[0], initials[1]
    }

    return initials[0], "_"
}

func main() {
    /*
    greeting("yoda")
    greeting("jedi")
    bye("obi-wan")

    cycleNames([]string{"obi-wan", "yoda", "kenobi"}, greeting)
    
    cycleNames([]string{"obi-wan", "yoda", "kenobi"}, bye)

    a1 := circleArea(10.5)
    a2 := circleArea(15)

    fmt.Printf("Area of circle 1 is %0.4f and circle 2 is %0.4f \n", a1, a2)
    */
    fn1, sn1 :=getInitials("obiwan kenobi")

    fmt.Println(fn1, sn1)
}
