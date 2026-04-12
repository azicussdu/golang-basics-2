package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("{Name: %s | Age: %d}\n", p.Name, p.Age)
}

//func main() {
//	prs := Person{Name: "Arman", Age: 21}
//
//	fmt.Println(prs) // {Name: Arman | Age: 21}
//}
