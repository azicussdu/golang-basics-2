package models

import "fmt"

var Num1 int
var num2 int

func Add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

type Person struct {
	Name string
	age  int
}

func (p Person) info() {
	fmt.Printf("Name: %s, age: %d\n", p.Name, p.age)
}
func (p Person) Show() {
	fmt.Printf("Name: %s, age: %d\n", p.Name, p.age)
}
