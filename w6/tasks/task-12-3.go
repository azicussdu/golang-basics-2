package main

import "fmt"

type Student struct {
	Name   string
	Grades []int // [80, 70, 50, ...]
}

func NewStudent(name string) *Student {
	st := Student{
		Name:   name,
		Grades: []int{},
	}
	return &st
}

func (st *Student) AddGrade(grade int) {
	if grade >= 0 && grade <= 100 {
		st.Grades = append(st.Grades, grade)
	}
}

func (st Student) ShowGrades() {
	fmt.Printf("%s grades: %v\n", st.Name, st.Grades)
}

func (st Student) Average() float64 {
	if len(st.Grades) == 0 {
		return 0
	}

	sum := 0
	for _, gr := range st.Grades {
		sum += gr // sum = sum + gr
	}

	return float64(sum) / float64(len(st.Grades))
}

//func main() {
//	st1 := NewStudent("Beka") // [70, 90]
//	st1.AddGrade(70)
//	st1.AddGrade(71)
//	st1.AddGrade(71)
//	fmt.Printf("%.3f\n", st1.Average())
//
//	st1.ShowGrades()
//}
