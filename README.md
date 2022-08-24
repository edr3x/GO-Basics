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
