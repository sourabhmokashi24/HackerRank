package main
import "fmt"

func ca(arr []int, n, k int) []int {
    var in = 0
    narr := make([]int, n)
    for j := k; j < n; j++ {
        narr[in] = arr[j]
        in++
    }
    for j := 0; j < k; j++ {
        narr[in] = arr[j]
        in++
    }
    return narr
}

func main() {
    var n, k, q int
    fmt.Scan(&n, &k, &q)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    k = (n - (k % n)) % n
    arr = ca(arr, n, k)
    var qr int
    for i := 0; i < q; i++ {
        fmt.Scan(&qr)
        fmt.Println(arr[qr])
    }
}
