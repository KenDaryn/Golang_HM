package HW_L_3

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


