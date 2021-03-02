//package main
//
//import "fmt"
//
////Задача №1 - калькулятор
////func main() {
////	for true {
////		var simvol string
////		fmt.Scanf("%s", &simvol)
////		if simvol == "0" {
////			fmt.Println("Конец")
////			break
////		}
////		var a, b int
////		fmt.Scanf("%d", &a)
////		fmt.Scanf("%d", &b)
////		if simvol == "+" {
////			fmt.Println("Итого по символу +: ", a+b)
////		} else if simvol == "*" {
////			fmt.Println("Итого по символу *: ", a*b)
////		} else if simvol == "/" {
////			fmt.Println("Итого по символу /: ", a/b)
////		} else if simvol == "-" {
////			fmt.Println("Итого по символу -: ", a-b)
////		} else {
////			fmt.Println("Вы ввели не правильный символ.")
////		}
////	}
////}
//
////Задача 2 - фибоначи
////func main() {
////		var number int
////		fmt.Scanf("%d", &number)
////		f1:=1
////		f2:=0
////		for i:=0; i<=number;i++{
////			f1,f2=f2,f1+f2
////			fmt.Println(f1)
////		}
////	}
//
////Задача №3 счетчик
////func main() {
////	var number int
////	fmt.Scanf("%d", &number)
////	shetchik := 0
////	for i := 1; i <= number; i++ {
////		var x int
////		fmt.Scanf("%d", &x)
////		shetchik = x + shetchik
////	}
////	fmt.Println("Итоговая сумма = ", shetchik)
////}
//
////func main () {
////var a, b int // за чем с и d? тут одно значение должно быть. Убрал переменную c и d.
////fmt.Scanf("%d", &a)
////x:=0 //новая переменная в которой будет хранится вводимые значение и сумирование.
//////i:=0 поменял на i:=1 так как цикл перебирает с цифры 0 из за этого перебирает 4 раза если вводить цифру 3
////for i := 1; i <= a; i++ {
////fmt.Scanf("%d", &b) //переменная b будет отображаться столько раз сколько равна числу a
////
////x=b+x // получается сумируешь x переменную которая равна 0 и переменную b которая равно вводимому значению.
////
//////за чем переменную d заменять переменной с а также зачем тут переменную b сумировать c
//////c = b
//////d = b + c
////}
//////fmt.Println(d)
////fmt.Println(x) //вывод итоговой суммы по переменной x
////}
//
////func main() {
////	//a := "" // эта переменная не нужна
////	var b, c int
////	var a string //надо добавить а как string
////	for true {
////		fmt.Scanf("%s", &a)
////		fmt.Printf(" Конец %s \n", a) //заменил текст на конец
////		if a == "0" {
////			break
////		}
////
////		if a == "+" || a == "-" || a == "/" || a == "*" {
////			fmt.Scanf("%d %d", &b, &c)
////			fmt.Printf("первое число: %d \n", b)
////			fmt.Printf("второе число: %d \n", c)
////			if a == "+" {
////				fmt.Printf("Ответ: %d \n", b+c)//сюда надо добавить \n что бы начиналось с новой строки что бы сумировать
////			}
////			if a == "-" {
////				fmt.Printf("Ответ: %d \n", b-c)
////			}
////			if a == "/" {
////				if c != 0 {
////					fmt.Printf("Ответ: %d \n", b/c)
////				} else {
////					fmt.Println("Деление на ноль!")
////				}
////			}
////			if a == "*" {
////				fmt.Printf("Ответ: %d \n", b*c)
////			}
////		}
////
////	}
////}
//
////func main() {
////	arr:=[]int{1,2,33,4,5}
////	println(arr[44])
////}
//
////func main() {
////	arr := [3]int{1, 2, 3}
////	fmt.Println(arr[0] + arr[1] + arr[2])
////	fmt.Println(len(arr))
////	var n int
////	arrq := [100]int{}
////	for i := 0; i < n; i++ {
////		var a int
////		fmt.Scanf("%d", &a)
////		arr[i]=a
////	}
////	z := 0
////	for i := 0; i < len(arr); i++ {
////		fmt.Println(arr[i])
////		z = z + arr[i]
////	}
////	fmt.Println(z)
////}
//
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i]=a
//	}
//	mini := 99999
//
//	for i := 0; i < n; i++ {
//		if arr[i] < mini {
//			mini = arr[i]
//		}
//	}
//	fmt.Println(mini)
//}
