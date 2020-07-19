package main
import "fmt"

func search(a int) int {
    var s int
    for i := 0; i <= a; i++ {
        if i % 2 == 0 {
            s += 1
        } else {
            s = s * 2
        }
    }
    return s
}

func main() {
    var n int
    fmt.Scan(&n)
    var s, a int
    for i := 0; i < n; i++ {
        fmt.Scan(&a)
        s = search(a)
        fmt.Println(s)
    }
}
