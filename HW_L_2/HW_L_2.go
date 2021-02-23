package main

//func main() {
//	var a int
//	fmt.Scanf("%d", &a)
//	c:=a%10
//	b:=a/10
//	v:=b%10
//	d:=a/100
//	println("Итого: ",c+v+d,a)
//}

//func main() {
//	var a, b int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	c := math.Pow(float64(a), 2.0) + math.Pow(float64(b), 2.0)
//	fmt.Println("Гипотенуза треугольника:", c)
//}

//y = kx + b (k-?,b-?)
//Ввод
//3 4
//5 6
//x1=3,y1=4
//x2=5,y2=6

//import "fmt"
//func main() {
//	// y = kx + b
//	var x1, y1, x2, y2 int
//	fmt.Scanf("%d", &x1)
//	fmt.Scanf("%d", &y1)
//	fmt.Scanf("%d", &x2)
//	fmt.Scanf("%d", &y2)
//	k := (y1 - y2) / (x1 - x2)
//	b := y2 - k * x2
//	fmt.Println(k,b)
//}

//func main(){
//	a:=0
//	b:=4
//	if a>b {
//		println("true")
//	}	else {
//		println("Falce")
//		}
//	}

//func  main()  {
//	var a,b int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	c:=a+b
//	if c>10{
//		fmt.Println(">10")
//	} else {
//		fmt.Println("a<10")
//	}
//}

//func  main()  {
//	var salary,month, price int
//	fmt.Scanf("%d", &salary)
//	fmt.Scanf("%d", &month)
//	fmt.Scanf("%d", &price)
//	r:=float64(price/month)
//	s:=float64(salary)*0.5
//
//	if s > r {
//		fmt.Println("Yes")
//	} else {
//		fmt.Println("No")
//	}
//}

//func main () {
//	var x1,y1,x2,y2,r int
//	fmt.Scanf("%d", &x1)
//	fmt.Scanf("%d", &y1)
//	fmt.Scanf("%d", &x2)
//	fmt.Scanf("%d", &y2)
//	fmt.Scanf("%d", &r)
//	x3:=x2-x1
//	y3:=y2-y1
//	S:=math.Sqrt(math.Pow(float64(x3),2)+math.Pow(float64(y3),2))
//	if S > float64(r) {
//		fmt.Println("No")
//	} else{
//		fmt.Println("Yes")
//	}
//}

//func main() {
//	a:=5
//	b:=6
//	c:=7
//	if a>b && a>c {
//		println("a is max")
//	} else {
//		println("a is not")
//	}
//}
//&& - and
//|| - или

//func main(){
//	var a,b,c int
//	fmt.Scanf("%d",&a)
//	fmt.Scanf("%d",&b)
//	fmt.Scanf("%d",&c)
//	if a+b<c || b+c<a || c+a<b {
//		fmt.Println("No")
//	} else {
//		fmt.Println("Yes")
//	}
//}

//1) висакосный год
//вводится одно число необходимо ответить является ли этот год висакосным или нет
//если не делится на 4
//	обычный
//если делится на 100
//	если делится на 400 - висакосный
//если не делится на 100 - вискаосный
//
//Год високосный, если он делится на четыре без остатка, но если он делится на 100 без остатка, это не високосный год.
//	Однако, если он делится без остатка на 400, это високосный год.
//	Таким образом, 2000 г. является особым високосным годом, который бывает лишь раз в 400 лет.

//func main() {
//	var year int
//	fmt.Scanf("%d", &year)
//	if year % 4 == 0 {
//		if year % 100 == 0 {
//			if year % 400 == 0 {
//				fmt.Println("Висакосный год.")
//			} else {
//				fmt.Println("Не висакосный год.")
//			}
//		} else {
//			fmt.Println("Висакосный год.")
//		}
//	} else {
//		fmt.Println("Не висакосный год.")
//	}
//}

//2) треугольник
//дается три стороны треугольник a,b, c
//необходимо понять сущестует ли треугольник и опрелить его тип
//равнобедренный, разносторний, равносторниий

//Разносторонним называется треугольник, у которого все три стороны не равны.
//Равнобедренным называется треугольник, у которого две стороны равны. Эти стороны называются боковыми, третья сторона называется основанием.
//	В равнобедренном треугольнике углы при основании равны.
//Равносторонним или правильным называется треугольник, у которого все три стороны равны.
//	В равностороннем треугольнике все углы равны 60°, а центры вписанной и описанной окружностей совпадают.
//	Равносторонний треугольник является частным случаем равнобедренного треугольника.

//func main(){
//	var a,b,c int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	fmt.Scanf("%d", &c)
//	if a==b && b==c {
//		fmt.Println("Равносторонний")
//	} else if a==b || a==c || b==c {
//		fmt.Println("Равнобедренний")
//	}else {
//		fmt.Println("Разносторний")
//	}
//}

//3) поиcк среднего числа
//среднее число то число которое больше первого но меньше второго
//ввод: a,b,c

//func main() {
//	var a, b, c int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	fmt.Scanf("%d", &c)
//	if a < b && b < c || c < b && b < a  {
//		fmt.Println("Среднее = b = ", b)
//	} else if b < a && a < c || c < a && a < b || b == c && c > a {
//		fmt.Println("Среднее = a = ", a)
//	} else if b < c && c < a   {
//		fmt.Println("Среднее = c = ", c)
//	} else if a == b && b < c || b == c && c < a || c == a && c < b || c == a && c > b || a == b && b > c || a < c && c < b{
//		fmt.Println("Два числа ровны")
//	} else {
//		fmt.Println("Все цифры ровны")
//	}
//}

//4) вводится одно число a и необходимо проверить является ли оно четным или нет

//func main () {
//	var number int
//	fmt.Scanf("%d", &number)
//	if number%2==0 {
//		fmt.Println("Четное число")
//	}else {
//		fmt.Println("Нечетное число")
//	}
//}