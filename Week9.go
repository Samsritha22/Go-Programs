package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter number of elements: ")
    fmt.Scan(&n)
    arr := make([]float64,n)
    sum := 0.0
    fmt.Println("Enter the elements:")
    for i:=0;i<n;i++ {
        fmt.Scan(&arr[i])
        sum += arr[i]
    }
    average := sum/float64(n)
    fmt.Println("Average=",average)
}
