package main

import "fmt"

//import "fmt"
//import "fmt"
//
//func main() {
//	var number int
//	fmt.Scanf("%d", &number)
//	if number <= 100 {
//		if number == 12 || number == 1 || number == 2 {
//			fmt.Println("Winter")
//		} else if number == 3 || number == 4 || number == 5 {
//			fmt.Println("Spring")
//		} else if number == 6 || number == 7 || number == 8 {
//			fmt.Println("Summer")
//		} else if number == 9 || number == 10 || number == 11 {
//			fmt.Println("Autumn")
//		} else {
//			fmt.Println("Error")
//		}
//	}else {
//		fmt.Println("Число не в диапазоне")
//	}
//}

//package main
//import (
//	"fmt"
//	"math"
//)
//func main() {
//	var q, w, e int
//	fmt.Scanf("%d", &q)
//	fmt.Scanf("%d", &w)
//	fmt.Scanf("%d", &e)
//	if float64(q) < math.Pow(10, 5) || float64(w) < math.Pow(10, 5) || float64(e) < math.Pow(10, 5) {
//		if q > w && w > e {
//			fmt.Println(q - e)
//		} else if w > q && q > e {
//			fmt.Println(w - e)
//		} else if q > e && e > w {
//			fmt.Println(q - w)
//		} else if w > e && e > q {
//			fmt.Println(w - q)
//		} else if e > w && w > q {
//			fmt.Println(e - q)
//		} else if e > q && q > w {
//			fmt.Println(e - w)
//		}
//	}
//}

//package main
//import (
//	"fmt"
//	"math"
//)
//func main() {
//	var q, w, e int
//	fmt.Scanf("%d", &q)
//	fmt.Scanf("%d", &w)
//	fmt.Scanf("%d", &e)
//	if float64(q) < math.Pow(10, 5) || float64(w) < math.Pow(10, 5) || float64(e) < math.Pow(10, 5) {
//		if q > w && w > e {
//			fmt.Println(float64(q) - float64(e))
//		} else if w > q && q > e {
//			fmt.Println(float64(w) - float64(e))
//		} else if q > e && e > w {
//			fmt.Println(float64(q) - float64(w))
//		} else if w > e && e > q {
//			fmt.Println(float64(w) - float64(q))
//		} else if e > w && w > q {
//			fmt.Println(float64(e) - float64(q))
//		} else if e > q && q > w {
//			fmt.Println(float64(e) - float64(w))
//		}
//	}
//}

//package main
//import (
//	"fmt"
//	"math"
//)
//func main(){
//	var n int
//	fmt.Scanf("%d",&n)
//	if float64(n) >=1 && float64(n)<=math.Pow(10,10000){
//		if n%11==0{
//			fmt.Println("YES")
//		} else {
//			fmt.Println("NO")
//		}
//	}
//}
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var H, A, B int
//	fmt.Scanf("%d", &H)
//	fmt.Scanf("%d", &A)
//	fmt.Scanf("%d", &B)
//	sumiA := 0
//	chetcik := 0
//	if H <= 1000 && B < A && A <= 100 {
//		for i := 0; i < H; i++ {
//			sumiA = sumiA + A - B
//			chetcik += 1
//			if sumi == 10 {
//				break
//			}
//		}
//	}
//	fmt.Println(sumi)
//	fmt.Println(chetcik)
//}


//struct
type Student struct {
	FirstName string
	LastName  string
	Age       int
	FullName  string
	YearBirth int
}

//methods
func (st *Student) PrintAll() {
	//create call of method CreateFullName()
	st.CreateFullName()
	st.SetYear()
	//create call of method SetYear()
	fmt.Println(st.FirstName, st.LastName, st.Age, st.FullName)
}



func (st *Student) CreateFullName() {
	st.FullName = st.FirstName + " " + st.LastName
}

//SetYear() method for calculating year of birth (2021-Age)


func (st*Student) SetYear(){
	st.YearBirth=2021-st.Age
}

func main() {
	var st1 Student
	st1 = Student{
		FirstName: "Yerassyl",
		LastName:  "Tleugazy",
		Age:       23,
	}
	st2 := Student{
		FirstName: "123231",
		LastName:  "4123313",
		Age:       123,
	}
	students := []Student{st1, st2}
	for _, v := range students {
		v.PrintAll()
	}
}





















