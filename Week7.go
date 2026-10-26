package main
import (
    "fmt"
    "strings"
)
func main() {
    var str string
    var left, right int
    fmt.Print("Enter a string:")
    fmt.Scan(&str)
    n := len(str)
    str = strings.ToLower(str)
    left = 0
    right = n-1
    palindrome := true
    for left <= right { 
        if str[left]!=str[right] {
            palindrome=false
            break
        }
        left++
        right--
    }
    if palindrome {
        fmt.Println(str,"is a palindrome")
    } else {
        fmt.Println(str,"is not a palindrome")
    }
}