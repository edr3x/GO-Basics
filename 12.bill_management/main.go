package main

import "fmt"

func main() {
    mybill := newBill("yoda's bill")

    mybill.addItem("soup", 4.50) 
    mybill.addItem("pie", 9.10)
    mybill.addItem("tandoori chicken", 44.2)
    mybill.addItem("Biryani", 7.50)

    mybill.updateTip(10)

    fmt.Println(mybill.format())
    
}
