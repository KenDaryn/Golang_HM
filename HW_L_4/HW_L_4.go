package main

//Задача №1 - калькулятор
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

//Задача 2 - фибоначи
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

//Задача №3 счетчик
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