//package main

//1)создать структура Book
//поля: name(string),
//цена(float)
//
//2) создать две переменные c типом Book
//3) посметить их в массив
//4) вывести элементы массива(процедура)
//Название Книги: книга1, Цена: 30.5

//func FullBoks(s Book) {
//	fmt.Println(s.Name, s.Price)
//}
//type Book struct {
//	Name  string
//	Price float64
//}
//func main() {
//	b1 := Book{
//		Name: "Книга №1",
//		Price:  1200.1,
//	}
//	b2 := Book{
//		Name:  "Книга №2",
//		Price: 9999.99,
//	}
//	students := []Book{b1, b2}
//	for _, v := range students {
//		FullBoks(v)
//	}
//}

//2) создать структура
//Студенты
//Строка FirstName
//Строка фамилии
//marks []int

//type Students struct {
//	Firstname string
//	Lastname  string
//	Marks     []int
//}
//
////строка getFullName()
//func GetFullName(s Students) {
//	fmt.Println(s.Firstname, s.Lastname)
//}
//
////getAverageMark() int
//func GetAverageMark(b Students) {
//	mrk := b.Marks
//	sumi := 0
//	for i := 0; i < len(mrk); i++ {
//		sumi += b.Marks[i]
//	}
//	avg := sumi / len(mrk)
//	fmt.Println(avg)
//}
//
//func main() {
//	s1 := Students{
//		Firstname: "Daryn",
//		Lastname:  "Kenzhebekov",
//		Marks:     []int{1, 2, 3, 4, 5},
//	}
//	s2 := Students{
//		Firstname: "Zhomart",
//		Lastname:  "Zhanuzakov",
//		Marks:     []int{3, 4, 1, 6, 7},
//	}
//	s3 := Students{
//		Firstname: "Marat",
//		Lastname:  "Ahmetov",
//		Marks:     []int{3, 4, 2, 6, 9},
//	}
//	student := []Students{s1, s2, s3}
//	for _, v := range student {
//		GetFullName(v)
//		GetAverageMark(v)
//	}
}

//2) создать структура
//Студенты
//Строка FirstName
//Строка фамилии
//marks []int


