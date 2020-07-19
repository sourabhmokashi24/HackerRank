package main
import "fmt"

func main() {
    var s int
    fmt.Scan(&s)
    for i := 0; i< s; i++ {
        var n, k int
        var c, a int
        fmt.Scan(&n, &k)
        for j := 0; j < n; j++ {
            fmt.Scan(&a)
            if a <= 0 {
                c++
            }
        }
        if c >= k {
            fmt.Println("NO")
        } else {
            fmt.Println("YES")
        }
    }
}
