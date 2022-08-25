# Bill Management

## What will we learn 

## Structs and Custom Types

```go
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

```
## Receiver Function

- There are functions which can be only used with predefined types

```go
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
```

> Note:
>
> By only changing the type of receiver function to pointer, GO automatically handles all other things inside functions
> We don't have to make variables inside the functions as pointer too.

## main.go file for the files above

```go
func main() {
    mybill := newBill("yoda's bill")

    mybill.addItem("soup", 4.50) 
    mybill.addItem("pie", 9.10)
    mybill.addItem("tandoori chicken", 44.2)
    mybill.addItem("Biryani", 7.50)

    mybill.updateTip(10)

    fmt.Println(mybill.format())
    
}
```

## User Input

- Taking user input from terminal insted of hardcoding
