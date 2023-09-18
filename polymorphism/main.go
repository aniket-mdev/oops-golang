package main

import "fmt"

// Polymorphism is the ability of any data to be processed in more than one form.
// The word itself indicates the meaning as poly means many and morphism means types.
// Polymorphism is one of the most important concepts of object-oriented programming languages.

// In Go, polymorphism can be achieved using interfaces.
// An interface is a collection of method signatures that any type can implement.
// This means that you can write code that can work with any type that implements an interface,
// without knowing the underlying type.

type Animal interface {
	Speak()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("\n%s  is speaking like mau mau ", c.Name)
}

func (d Dog) Speak() {
	fmt.Printf("\n%s is speaking like bho bho ", d.Name)
}

func printAnimals(a Animal) {
	a.Speak()
}

func main() {
	c := Cat{"cat"}
	d := Dog{"dog"}

	printAnimals(c)
	printAnimals(d)
}
