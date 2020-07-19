package main
import "fmt"

func cal(a int) int {
    var c int  
    tmp := a
    for tmp > 0 {
        r := tmp % 10
        if r != 0 && (a % r == 0) {
            c++
        }
        tmp = tmp / 10
    }
    return c
}

func main() {
    var n int
    fmt.Scan(&n)
    var a, s int
    for i := 0; i < n; i++ {
        fmt.Scan(&a)
        s = cal(a)
        fmt.Println(s)
    }
}
