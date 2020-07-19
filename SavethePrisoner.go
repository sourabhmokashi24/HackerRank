package main
import "fmt"

func main() {
    var s2 int
    fmt.Scan(&s2)
    for i := 0; i < s2; i++{
        var n, m, s, s1 int
        fmt.Scan(&n, &m, &s)
        s1 = m % n
        s1 = s1 + (s - 1)
        s1 = s1 % n
        if s1 == 0 {
            s1 = n
        }
        fmt.Println(s1)
    }
}
