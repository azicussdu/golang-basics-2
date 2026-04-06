package main

//import "fmt"
//
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//type Circle struct {
//	Radius float64
//}
//
//func (cr *Circle) Area() float64 {
//	return 3.14 * (cr.Radius * cr.Radius)
//}
//
//func (cr *Circle) Perimeter() float64 {
//	return 2 * 3.14 * cr.Radius
//}
//
//func shapeInfo(sh Shape) { // sh = *Circle (tak kak recievery - eto *Circle)
//	fmt.Printf("Type: %T\n", sh)
//	fmt.Println("Shape area: ", sh.Area())
//	fmt.Println("Shape perimeter: ", sh.Perimeter())
//}

//func main() {

//cir := Circle{Radius: 5.5}
//shapeInfo(cir) // Oshibka (tek *Circle realizuet Shape)

//cir2 := &Circle{Radius: 5.5} //   Circle - Not OK   VS   *Circle - OK
//shapeInfo(cir2)

// var _ Shape = Circle{}       // Oshibka (tak kak ne Circle a *Circle realizuet Shape)
//var _ Shape = (*Circle)(nil) // var ptr *Circle = nil (obekt kurmai tekseru)

//}
