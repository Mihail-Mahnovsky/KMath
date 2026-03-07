# KMath is simple light weight mathlib 

examples:
```go
package main

import (
    "fmt"
    "github.com/Mihail-Mahnovsky/linalg-go"
)

func main() {
    v1 := linalg.MakeVec3(1, 2, 3)
    v2 := linalg.MakeVec3(4, 5, 6)
    sum := v1.Add(v2)
    sum.Print()
}
```