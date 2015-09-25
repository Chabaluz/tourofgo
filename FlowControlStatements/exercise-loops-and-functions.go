package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    y := 0.0
    for i := 0; z - y > 0.01 || z - y < -0.01; i++ {
        y = z
        z = z -((math.Pow(z, 2) - x) / 2 * z)
        fmt.Println(i, y, z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println("math.Sqrt(2) = ", math.Sqrt(2))
}
