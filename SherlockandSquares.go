package main
import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scan(&n)
    for j := 0; j < n; j++ {
        var a, b, c, s int
        fmt.Scan(&a, &b)
        for i := int(math.Sqrt(float64(a))); i <= int(math.Sqrt(float64(b))); i++ {
            s = i * i
            if s >= a && s <= b {
                c++
            }
        }
        fmt.Println(c)
    }
}
