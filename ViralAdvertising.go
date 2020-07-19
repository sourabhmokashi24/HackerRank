package main
import "fmt"

func main() {
    var n, s, l int
    s = 5
    l = 2
    fmt.Scan(&n)
    for i := 1; i < n; i++ {
        s = (s / 2) * 3
        l += (s / 2)
    }
    fmt.Println(l)
}
