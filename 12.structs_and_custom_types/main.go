package main

import "fmt"

func main() {
    mybill := newBill("yoda's bill")

    fmt.Println(mybill.format())
    
}
