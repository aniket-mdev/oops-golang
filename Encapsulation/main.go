package main

import (
	"fmt"

	student "github.com/aniket0951/OOP/Encapsulation/Student"
)

// binding variables and methods together into a single unit
// and preventing them from being accessed by other classes

// Encapsulation is achieved by exported elements(variables, functions, methods, fields, structures) from the packages,
// it helps to control the visibility of the elements(variables, functions, methods, fields, structures).
// The elements are visible if the package in which they are defined is available in your program

func main() {
	s := student.Student{RollNo: 1}
	s.SetName("Aniket")
	s.SetCourse("Hacking")

	fmt.Println(s.GetName())
	fmt.Println(s.GetCourse())
}
