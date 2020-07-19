package main
import "fmt"

func find(arr []int, i, n int) int {
    var j int
    for j = 0; j < n; j++ {
        if arr[j] == i {
            break
        }
    }
    return j+1
}

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    var s int
    for i := 0; i < n; i++ {
        s = find(arr, i+1, n)
        s = find(arr, s, n)
        fmt.Println(s)
    }
}
