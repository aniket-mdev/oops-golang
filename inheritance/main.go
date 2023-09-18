package main

import "fmt"

/*

Golang does not support classes, so No Inheritance Concept in Golang.
But we can give a behaviour of Inheritance to the Program.
Inheritance takes place through struct embedding.
We cannot directly extend structs but rather use a concept called composition,
where the struct is used to form other objects


*/
type Users struct {
	Name      string
	Age       int64
	ContactNo int64
}

type Student struct {
	Users
	RollNo int
}

func main() {
	// creating an instance
	s := Student{
		Users: Users{
			Name:      "XYZ",
			Age:       10,
			ContactNo: 1234567890,
		},
		RollNo: 1,
	}
	fmt.Printf("User data %+v: \n", s)
}
