package main

import "fmt"

func updateName(n string){
    n = "yoda"
}

func updateMenu(y map[string]float64){
    y["tea"] = 25.33
}

func main() {
    
    name:= "kenobi"

    updateName(name)

    fmt.Println(name)


    //group B types -> slices, maps, functions

    menu := map[string]float64{
        "momo": 120.5,
        "ice-cream": 30.5,
    }

    updateMenu(menu)

    fmt.Println(menu)

}
