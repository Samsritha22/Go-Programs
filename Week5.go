package main

import "fmt"

func main() {
    var n, i, j int

    fmt.Println("Enter the number of rows: ")
    fmt.Scan(&n)

    num := 1

    for i = 1; i <= n; i++ {
        for j = 1; j <= i; j++ {
            fmt.Print(num, " ")
            num++
        }
        fmt.Println()
    }
}