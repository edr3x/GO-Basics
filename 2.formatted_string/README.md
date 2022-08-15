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
