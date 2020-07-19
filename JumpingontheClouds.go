package main
import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    var i int = 0
    var s int = 100
    for {
        ti := (i + k) % n        
        if arr[ti] == 0 {
           s -= 1
        } else {
           s -= 3
        }
        i = ti
        if i == 0 {
            break
        }
    }
    fmt.Println(s)
}
