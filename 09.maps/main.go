package main

import "fmt"

func main() {
	menu := map[string]float64{
		"momo":    150.0,
		"noodles": 100.0,
		"soup":    35.5,
		"tea":     25.5,
	}

	fmt.Println(menu)
    fmt.Println(menu["soup"])

    // Looping through the map

    for k,v := range menu{
        fmt.Println(k, "-", v)
    }

    // interger key type 

    phonediary := map[int]string{
        98524: "Jhon",
        98443: "Doe",
        97469: "Kenobi",
    }

   fmt.Println(phonediary) 
   fmt.Println(phonediary[98524])

   phonediary[97469] = "yoda"

   fmt.Println(phonediary)
   
}
