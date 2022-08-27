# GO Learning

## Simple Hello World program

```go
package main

import "fmt"

func main() {
 fmt.Println("Hello World")
}
```

### Run the program

```bash
go run hello.go
```

# Variables

## String

- String is declared by `string` keyword

```go
var firstName string = "Obi-Wan"

var lastName = "kenobi"

var signature string

signature = "Hello There!"
```

- String Interpolation

```go
fmt.Println(firstName, lastName, "says", signature)
// prints "Obi-Wan kenobi says Hello There!"
```

- Variable Declaration Shorthand

```go
func main(){
    name := "Obi-Wan Kenobi"
}
```

> Note
>
> - Variable declaration shorthand can only be used inside a function.

## Integer

- Signed integer is declared by `int` keyword and `uint` for unsigned numbers

```go
// Signed integer
var randomOne int = -33
var randomTwo = 555
randomThree := int(-33) // Shorthand with type defined
randomThree := -49

// Unsigned integer
var ageOne int = 9
var ageTwo = 30
ageThree := 50
ageFour := uint(50) // Shorthand with type defined
```

- For bit wise integer operations, use `int8` or `int16` or `int32` or `int64`
- Also use `uint8` or `uint16` or `uint32` or `uint64` for unsigned integers

```go
var numOne int8 = -25
var numTwo uint8 = 44
numThree := int8(-25)
numFour := 88
```

## Float

- Float is declared by `float32` or `float64`
- If type is not defined then default type is `float64`

```go
var randomFloat1 float32 = -33.33
var randomFloat2 float64 = 323.33890798479102748912789347298
randomFloat3 := float32(-33.33)
randomFloat4 := 323.33890798479102748912789347298
```

# Formatted String

## Format Specifiers start with `%` and are followed by a format specifier

### Some format specifiers are

- v - default format specifier for values
- d - decimal integer
- o - octal integer
- O - octal integer with 0o prefix
- b - binary integer
- x - hexadecimal integer lowercase
- X - hexadecimal integer uppercase
- f - decimal floating point, lowercase
- F - decimal floating point, uppercase
- e - scientific notation (mantissa/exponent), lowercase
- E - scientific notation (mantissa/exponent), uppercase
- g - the shortest representation of %e or %f
- G - the shortest representation of %E or %F
- c - a character represented by the corresponding Unicode code point
- q - a quoted character
- U - Unicode escape sequence
- t - the word true or false
- s - a string
- #v - Go-syntax representation of the value
- T - a Go-syntax representation of the type of the value
- p - pointer address
- % - a double %% prints a single %

```go
name := "Yoda"
race := "Jedi"
age := 200

//* Printf (Formatted strings)
fmt.Printf("Name of mine is %v and im a %v \n", name, race) //* Default value
        // Name of mine is Yoda and im a Jedi
fmt.Printf("Name is %q and age is %d \n", name, age)        //* quoted value with decimal
        // Name is "Yoda" and age is 200
```

```go
//* Sprintf (Formatted strings)
savedStr := fmt.Sprintf("Name is %q and age is %d \n", name, age)
fmt.Println("the saved string is:", savedStr)
        // the saved string is: Name is "Yoda" and age is 200
```

> Note
>
> `Sprint` keyword is used to store strings or values in a variable

# Array & Slices

## Arrary

- declared by `var <varname> [<length>]<datatype> = [<length>]<datatype>{<values>}`

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5}

var arr = [5]int{1, 2, 3, 4, 5}


names := [2]string{"jhon", "doe"}

names[0] = "nojhon" // change value but can't change length

```

## Slices

- Array under the hood but without fixed length

```go
var scores = []int{445, 232, 435}

scores[0] = 670

scores = append(scores, 85) // adds 85 to end of slice
```

- Append() alone won't change the slice but returns a new slice so reassigned it to the original variable above

## Range Slices

```go
names := [4]string{"jhon", "doe", "obi-wan", "kenobi"}

range1 := names[1:3] // {doe, obi-wan}
range2 := names[2:] // {obi-wan, kenobi}
range3 := names[:3] // {jhon, doe, obi-wan}

fmt.Println(range1, "\n", range2, "\n", range3)

range1 = append(range1, "broda") // adds broda to end of range1

fmt.Println(range1)

```

# Loops

## While loop

```go
    x:= 0

    for x<5 { // while loop
        fmt.Println("value of x is:", x)
        x++
    }
```

## For loop

```go
    names := []string{"obi-wan","yoda", "anakin", "gervious", "leila"}

    for i := 0 ; i<len(names); i++{ // for loop
        fmt.Println(names[i])
    }
```

## For loop with range

- just like for in loop in JS

```go
    names := []string{"obi-wan","yoda", "anakin", "gervious", "leila"}

    for index, val := range names{
        fmt.Printf("the value at position %v is %v \n", index, val)
    }

    for _,val := range names{ // if we don't want index and only value then
        fmt.Printf("the value is %v \n", val)
    }

```

# Conditonals

```go
	age := 25

    if age < 30 {
        fmt.Println("is less than 30 ")
    } else if age< 40 {
        fmt.Println("age is less than 40")
    } else {
        fmt.Println("age is not less than 45")
    }

    names := []string{"obi-wan","yoda", "anakin", "gervious", "leila"}

    for index, val := range names{
        if index == 1 {
            fmt.Println("continuing at pos ", index)
            continue
        }

        if index > 2 {
            fmt.Println("breaking at pos", index)
            break
        }

        fmt.Printf("the value at pos %v is %v \n", index, val)
    }
```

# Functions

## normal functions

```go

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

func main() {
    cycleNames([]string{"obi-wan", "yoda", "kenobi"}, greeting)
    cycleNames([]string{"obi-wan", "yoda", "kenobi"}, bye)
}

```

## Functions with Returntype

- Returntype of function is define between () and {} of a function 
- Like in bellow example we have return type `float64` returntype

```go
func circleArea(r float64) float64{
    return math.Pi * r * r
}

func main(){
    a1 := circleArea(10.5)
    a2 := circleArea(15)

    fmt.Printf("Area of circle 1 is %0.4f and circle 2 is %0.4f \n",a1, a2)
}
```
### Multiple Returntype

```go

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

func main(){
    fn1, sn1 :=getInitials("obiwan kenobi")

    fmt.Println(fn1, sn1)
}
```

# Package Scope

- if files are in same package name i.e. `package main` for this example 
- they share the same scope and we can access variables across files if they are global variable

## Example

### main.go

```go
package main

import "fmt"

var order = 66

func main() {

	sayHello("kenobi")

	for _, v := range points {
		fmt.Println(v)
	}

    orderNum()
}
```

### greeting.go

```go
package main

import "fmt"

var points = []int{20, 55, 343, 23, 45}

func sayHello(n string) {
	fmt.Println("hello", n)
}

func orderNum(){
    fmt.Println("order number is :", order)
}
```

- As shown above we are accessing the variables across files 
- we have to run both files at same time to work 

```sh
go run main.go greeting.go
```

# Maps

- like other languages all of the keys can be of multiple different primitive type like strings, floats e.t.c.
- but unlike others all of the keys in single map must have same type 
- also all of the values of single map must have single type

```go
varname := map[string]float64{}
```

- here we created a map with key of type `string` and value of type `float64`

```go
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
 
```

> Note:
> 
> Key and value can be of any type but in single map there can be only one type of key
> and one type of value.

# Passing

- Go is a Pass-by-value language
- Go makes copies of values when passed into functions

```go
func updateName(n string){
    n = "yoda"
}

func main() {
    name:= "kenobi"

    updateName(name)

    fmt.Println(name)
}
```

- We get output `kenobi` because function passes the copy of the original value not the original value
- but if we want to change the original value then we can do this:

```go
func updateName(n string) string {
    n = "yoda"
    return x
}

func main() { 
    name:= "kenobi"

    name = updateName(name)

    fmt.Println(name)
}
```

- This behaviour is done by `strings, ints, bools, floats, arrays, structs`
- can be called as `Non-Pointer Values`

### now for another type

```go
func updateMenu(y map[string]float64){
    y["tea"] = 25.33
}

func main(){
    menu := map[string]float64{
        "momo": 120.5,
        "ice-cream": 30.5,
    }

    updateMenu(menu)

    fmt.Println(menu)
}
```

- Here we get updated menu that is we get additional key value pair on output
- This behaviour is done by `Slices, Maps functions`
- Can be called as `Pointer wrapper values`

# Pointer

- Pointer are are datatypes which points to another memory address 
- denoted by `&varname`
- can be deferenced by using `*`

```go
    name := "yoda"

    mem := &name

    fmt.Println("memory address is :", mem) // prints memory address
    fmt.Println("value at memory address is :", *mem) // prints 'yoda'
   
```

### another example 

```go
func main() {
    name := "yoda"

    mem := &name
    
    fmt.Println(name) // prints "yoda"
    updateName(mem)
    fmt.Println(name) // prints "kenobi"

}

func updateName(n *string){
    *n = "kenobi"
}
```

- Here we passed the memory address of variable `name` so the value pointed by function is also the same memory address
- Since original memory address is pointed thus value of variable `name` is changed

# Structs and Custom Types

```go
type bill struct{
    name string
    items map[string]float64
    tip float64
}

// make new bills
func newBill(name string) bill {
    b := bill{
        name:name,
        items: map[string]float64{},
        tip: 0,
    }

    return b
}

func main() {
    mybill := newBill("yoda's bill")

    fmt.Println(mybill)
    
}
```

- this way we can create a struct and use it

## Receiver Function in struct

- Receiver Function is like normal function but can only be used via specified struct 

```go
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//total
    fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total) // -25 gives 25 spaces on right
                                                        // +25 gives 25 spaces on left

    return fs
}
```
 
- in main.go

```go
func main() {
    mybill := newBill("yoda's bill")

    fmt.Println(mybill.format())
    
}
```

## Other things like Structs, Parsing and saving to file are [here](https://github.com/edr3x/GO-Learning/tree/master/12.bill_management)

## Interfaces

```go
type shape interface {
    area() float64
    circumf() float64
}
```

## Goroutine

- A goroutine is a lightweight thread managed by the Go runtime

```
    // normal sync function call
    foo(var)
    
    // goroutine
    go foo(var)
```
