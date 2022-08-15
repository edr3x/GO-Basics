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
