package main

//1) висакосный год

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

//func main()  {
//	for i:=0;i<5;i++{
//		fmt.Println("Daryn")
//	}
//}

//for создается счетчи; условие работы счетчика;момент изменения счетчика{дейсвтие 1, действие 2}

//Цикл - выгрузка нечетных чисел
//func main (){
//	for i:=0; i<=100;i++{
//		if i%4==0 {
//			fmt.Println(i,"")
//		}
//	}
//	for i:=0; i<=100; i+=4{
//		fmt.Println(i,"")
//	}
//}

//func main() {
//	for i:=0;i<50;i++{
//		if i%2==0{
//			fmt.Println(i)
//		}
//	}
//}

//func main() {
//	for i:=0;i<50;i+=2{
//			fmt.Println(i)
//		}
//	}

//func main(){
//	sumi:=0
//	for i:=0; i<5; i++{
//		sumi=sumi+1
//	}
//	fmt.Println(sumi)
//}

//func main(){
//	a:=0
//	for i:=0;i<20;i++{
//		a+=i
//	}
//	fmt.Println(a)
//}

//func main() {
//	var a,b int
//	fmt.Scanf("%d %d", &a,&b)
//	if a<b {
//		for i:=a; i<b;i++{
//			
//		}
//	}
//}

//Факториал
//func main(){
//	var a int
//	fmt.Scanf("%d", &a)
//	fac:=1
//	for i:=1;i<=a;i++{
//		fac=i*fac
//	}
//	fmt.Println(fac)
//}

//func main() {
//	for i:=0; i<=10; i++ {
//		if i == 5 {
//			//break - останавляет цикл
//			//continue - пропускает шаг и дальше продолжает цикл
//		}
//		fmt.Println(i)
//	}
//}

//func main() {
//	var a, b int
//	fmt.Scanf("%d %d", &a, &b)
//	chet := 0
//	nechet := 0
//	for i := a; i <= b; i++ {
//		if i%2 == 0 {
//			chet += 1
//		} else {
//			nechet += 1
//		}
//	}
//	fmt.Println("Четное: ", chet,"Не четное: ", nechet)
//}

//Задача №3
//func main() {
//	var number int
//	fmt.Scanf("%d", &number)
//	shetchik := 0
//	for i := 1; i <= number; i++ {
//		var x int
//		fmt.Scanf("%d", &x)
//		shetchik = x + shetchik
//	}
//	fmt.Println("Итоговая сумма = ", shetchik)
//}

//Задача №1
//func main() {
//	for true {
//		var simvol string
//		fmt.Scanf("%s", &simvol)
//		if simvol == "0" {
//			fmt.Println("Конец")
//			break
//		}
//		var a, b int
//		fmt.Scanf("%d", &a)
//		fmt.Scanf("%d", &b)
//		if simvol == "+" {
//			fmt.Println("Итого по символу +: ", a+b)
//		} else if simvol == "*" {
//			fmt.Println("Итого по символу *: ", a*b)
//		} else if simvol == "/" {
//			fmt.Println("Итого по символу /: ", a/b)
//		} else if simvol == "-" {
//			fmt.Println("Итого по символу -: ", a-b)
//		} else {
//			fmt.Println("Вы ввели не правильный символ.")
//		}
//	}
//}

//Задача 2
//func main() {
//		var number int
//		fmt.Scanf("%d", &number)
//		f1:=1
//		f2:=0
//		for i:=0; i<=number;i++{
//			f1,f2=f2,f1+f2
//			fmt.Println(f1)
//		}
//	}

//Задача 3
//func main() {
//	var number int
//	fmt.Scanf("%d", &number)
//	f1:=1
//	f2:=0
//	for i:=0; i<=number;i++{
//		//i:=f2
//
//		i=f1+f2
//		f1=f2
//		f2=i
//		fmt.Println(i)
//	}
//}
