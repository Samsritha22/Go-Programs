package main

import "fmt"

type Contact struct {
    Name    string
    Email   string
    Message string
}

func main() {
    var contact Contact

    fmt.Print("Enter your Name: ")
    fmt.Scan(&contact.Name)

    fmt.Print("Enter your Email: ")
    fmt.Scan(&contact.Email)

    fmt.Print("Enter your Message: ")
    fmt.Scan(&contact.Message)

    fmt.Println("\nContact Details:")
    fmt.Println("Name:", contact.Name)
    fmt.Println("Email:", contact.Email)
    fmt.Println("Message:", contact.Message)
}