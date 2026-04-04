package main

type Student struct {
	Name  string
	Grade string
}

func groupByGrade(students []Student) map[string][]string {
	mp := make(map[string][]string)

	for _, s := range students {
		mp[s.Grade] = append(mp[s.Grade], s.Name)
	}

	return mp
}

//func main() {
//	students := []Student{
//		{Name: "Arman", Grade: "A"},
//		{Name: "Beka", Grade: "B"},
//		{Name: "Seka", Grade: "A"},
//		{Name: "Zheka", Grade: "C"},
//		{Name: "Dake", Grade: "B"},
//	}
//	resultMap := groupByGrade(students)
//	fmt.Println(resultMap)
//}
