package main

//import "fmt"
//
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//type Square struct {
//	Side int
//}
//
//func (sq Square) Area() float64 {
//	return float64(sq.Side * sq.Side)
//}
//
//func (sq Square) Perimeter() float64 {
//	return float64(sq.Side * 4)
//}
//
//type Circle struct {
//	Radius float64
//}
//
//func (cr Circle) Area() float64 {
//	return 3.14 * (cr.Radius * cr.Radius)
//}
//
//func (cr Circle) Perimeter() float64 {
//	return 2 * 3.14 * cr.Radius
//}
//
//func shapeInfo(sh Shape) { // sh = Circle, *Circle (ekeui de Shape-ty realizuyut)
//	fmt.Printf("Type: %T\n", sh)
//	fmt.Println("Shape area: ", sh.Area())
//	fmt.Println("Shape perimeter: ", sh.Perimeter())
//}

//func main() {

//	cir := &Circle{Radius: 5.5} //   Circle   VS   *Circle
//	shapeInfo(cir)
//
//	var _ Shape = Circle{}       // tek tekseru ushin (no tekseru ushin prishlos sozdat obekt)
//	var _ Shape = (*Circle)(nil) // var ptr *Circle = nil (obekt kurmai tekseru)

//sq := Square{Side: 15}
//shapeInfo(sq)
//
//cir := Circle{Radius: 5.5}
//shapeInfo(cir)
//
//shapes := []Shape{sq, cir}
//
//for _, sh := range shapes {
//	shapeInfo(sh)
//}

//var x Shape
//x = Square{Side: 10}
//
//fmt.Println("Area: ", x.Area())
//fmt.Println("Perimeter: ", x.Perimeter())
//}
