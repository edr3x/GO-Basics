package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//Receiver function to format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

    //adding the tip
    fs += fmt.Sprintf("%-25v ...$%0.2f", "tip:", b.tip)

	//total
    fs += fmt.Sprintf("\n%-25v ...$%0.2f", "total:", total+b.tip)

    return fs
}

// Update the tip
func (b *bill) updateTip(tip float64){
    b.tip = tip
}

// Adding item to the map
func (b *bill) addItem(name string, price float64){
    b.items[name] = price
}
