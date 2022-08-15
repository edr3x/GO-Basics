# GO Learning

## Simple Hello World program

```go
package main

import "fmt"

func main() {
 fmt.Println("Hello World")
}
```

## Variables

### String

```go
var firstName string = "Obi-Wan"

var lastName = "kenobi"

var signature string

signature = "Hello There"
```

- String Interpolation

```go
fmt.Println(firstName, lastName, "says", signature)
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
