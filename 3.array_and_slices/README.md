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
