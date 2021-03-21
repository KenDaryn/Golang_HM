package main

import "fmt"

type Students struct {
	Firstname string
	Lastname  string
	Marks     []int
}

func (z Students) getFullName() {
	fmt.Println(z.Firstname, z.Lastname)
}

func (a Students) getAverageMark() {
	sumi := 0
	avg := 0
	for i := 0; i < len(a.Marks); i++ {
		sumi += a.Marks[i]
	}
	avg = sumi / len(a.Marks)
	fmt.Println(avg)
}

func (d Students) Fqwe() {
	student := []Students{}
	for _, v := range student {
		fmt.Println(v)

	}
}
func main() {
	s1 := Students{
		Firstname: "Daryn",
		Lastname:  "Kenzhebekov",
		Marks:     []int{1, 2, 3, 4, 5},
	}
	s2 := Students{
		Firstname: "Zhomart",
		Lastname:  "Zhanuzakov",
		Marks:     []int{3, 4, 1, 6, 7},
	}
	s3 := Students{
		Firstname: "Marat",
		Lastname:  "Ahmetov",
		Marks:     []int{3, 4, 2, 6, 9},
	}
	s1.getFullName()
	s1.getAverageMark()
	s2.getFullName()
	s2.getAverageMark()
	s3.getFullName()
	s3.getAverageMark()
}