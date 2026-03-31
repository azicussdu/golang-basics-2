package main

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	if width <= 0 {
		width = 1
	}
	if height <= 0 {
		height = 1
	}

	rec := Rectangle{Width: width, Height: height}
	return &rec
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) ChangeSides(width float64, height float64) {
	if width > 0 {
		r.Width = width
	}
	if height > 0 {
		r.Height = height
	}
}

//func main() {
//	rect := NewRectangle(10, 15)      // rect = &Rectangle{}
//	fmt.Println("area:", rect.Area()) // (*rect).Area()
//	fmt.Println("perimeter:", rect.Perimeter())
//
//	rect.ChangeSides(0, 40)
//	fmt.Println(rect)
//}
