package main

import (
	"fmt"
	"studentproject/student"
)

func main() {
	s := student.Student{
		Name:   "Samsritha",
		Age:    21,
		Rollno: 33,
	}

	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Roll no:", s.Rollno)
}
