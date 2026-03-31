package main

import "fmt"

type BaseModel struct {
	ID        int
	CreatedAt string
}

func (b BaseModel) Info() string {
	str := fmt.Sprintf("ID: %d - Created at: %s\n", b.ID, b.CreatedAt)
	return str
}

type Product struct {
	BaseModel
	Name  string
	Price float64
}

func (pr Product) Info() string {
	str := fmt.Sprintf("ID: %d - Created at: %s, Name: %s and Price: %.2f\n", pr.ID, pr.CreatedAt, pr.Name, pr.Price)
	return str
}

//func main() {
//	pr := Product{}
//	pr.ID = 12
//	pr.CreatedAt = "22-03-2022"
//	pr.Name = "Iphone"
//	pr.Price = 230000
//	fmt.Print(pr.Info())
//
//	bm := BaseModel{11, "10-01-2026"}
//	fmt.Print(bm.Info())
//}
