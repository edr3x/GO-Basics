package main

import "fmt"

func main() {
    name := "yoda"

    mem := &name

    fmt.Println("memory address is :", mem)
    fmt.Println("value at memory address is :", *mem)
    
    fmt.Println(name)
    updateName(mem)
    fmt.Println(name)

}

func updateName(n *string){
    *n = "kenobi"
}
