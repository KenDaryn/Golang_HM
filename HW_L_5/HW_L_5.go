package main

//1) найти количество четных и нечетных элементов в массиве - готово
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	chet := 0
//	nechet := 0
//	for i := 0; i < n; i++ {
//		x := arr[i]
//		if x%2 == 0 {
//			chet += 1
//		} else {
//			nechet += 1
//		}
//	}
//	fmt.Println("Четное = ", chet, "Нечетное = ", nechet)
//}

//2) найти сумму положительных элементов в массиве
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//		z := -1
//	for i := 0; i < n; i++ {
//		if arr[i] > z {
//			fmt.Println(arr[i])
//		}
//	}
//}

//3) найти количество элементов которые меньше среднего арифметического элементов и вывыести их
//среднееф ариф avg = sumi/n

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	summ := 0
//	for i := 0; i < n; i++ {
//		summ = summ + arr[i]
//	}
//	n1 := 0
//	avg := 0
//	for i := 0; i < n; i++ {
//		if summ/n > arr[i] {
//			fmt.Println("Меньше среднего", arr[i])
//			avg = avg + arr[i]
//			n1 += 1
//		}
//	}
//	fmt.Println("Среднее чсило : ", avg/n1)
//}

//4) найти разницу между максимальным и минимальными элементами массива

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	max := 0
//	mini := 9999
//	for i := 0; i < n; i++ {
//		if arr[i] > max {
//			max = arr[i]
//		} else if arr[i] < mini {
//			mini = arr[i]
//		}
//	}
//	fmt.Println("Максимальная сумма: ", max)
//	fmt.Println("Минимальная сумма", mini)
//	fmt.Println("Разница между max и mini: ", max-mini)
//}
//5) найти в массиве четные и нечетные элементы и
//поместить их в разные массивы и вывести эти массивы
//
//func main(){
//		var n int
//		fmt.Scanf("%d", &n)
//		arr := [100] int{}
//		for i := 0; i < n; i++ {
//			var a int
//			fmt.Scanf("%d", &a)
//			arr[i] = a
//}
//	chet := [10] int{}
//	nechet:=[10] int{}
//	for i := 0; i < n; i++ {
//		if arr[i]%2==0{
//			chet[i]=arr[i]
//		} else{
//			nechet[i]=arr[i]
//		}
//	}
//	fmt.Println("Четные = ", chet)
//	fmt.Println("Не четные", nechet)
//}

//1) Вводятся три разных числа. Найти, какое из них является средним (больше одного, но меньше другого).

//func main() {
//	var a, b, c int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	fmt.Scanf("%d", &c)
//
//	if a < b && b < c || c < b && b < a {
//		fmt.Println("b", b)
//	} else if b < a && a < c || c < a && a < b {
//		fmt.Println("a", a)
//	} else if b < c && c < a || a < c && c < b {
//		fmt.Println("c", c)
//	} else {
//		fmt.Println("a,b,c", a, b, c)
//	}
//}

//2) Из двух случайных чисел, одно из которых четное, а другое нечетное, определить и вывести на экран нечетное число.

//func main() {
//	var a, b int
//	fmt.Scanf("%d", &a)
//	fmt.Scanf("%d", &b)
//	if a%2 != 0 && b%2 != 0 {
//		fmt.Println("Оба нечетные", a, b)
//	} else if b%2 != 0 {
//		fmt.Println(b)
//	} else if a%2 != 0 {
//		fmt.Println(a)
//	} else {
//		fmt.Println("Все четные", a, b)
//	}
//}

//3) В массиве найти максимальный элемент с четным индексом.
//Другая формулировка задачи: среди элементов массива с четными индексами, найти тот, который имеет максимальное значение.

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [100]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	max := 0
//	for i := 0; i < n; i++ {
//		if arr[i]%2 == 0 {
//			if arr[i] > max {
//				max = arr[i]
//			}
//		}
//	}
//	fmt.Println("Максимальная четная сумма: ", max)
//}

//4) Вывести элементы заданного массива в обратном порядке, то есть произвести реверс массива.

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [5]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//
//	for i := n; i <= n; i-- {
//		if i==-1{
//			break
//		}
//		fmt.Println(arr[i])
//
//	}
//}


//5) Все элементы массива поделить на значение наибольшего элемента этого массива.
//И сохранить в текущий массив

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [5]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	max := 0
//	for i := 0; i < n; i++ {
//		if arr[i] > max {
//			max = arr[i]
//		}
//	}
//	for i := 0; i < n; i++ {
//		arr[i] = arr[i] / max
//	}
//	fmt.Println(arr)
//}

//1.Напишите формулу чтобы вычислять последний индекс массива, формула должна работать для всех массивов
//
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [5]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	fmt.Println(len(arr)-1)
//}
//


//package main

import "fmt"

func GetB(c []int) int {
	a := 0
	b := 0
	for i := 0; i < len(c); i++ {
		if c[i] < 1 {
			a = 0
		} else {
			a = a + 1
		}
		if a > b {
			b = a
		}
	}
	return b
}
func printB(d []int) {
	number := GetB(d)
	fmt.Println("Длина самой продолжительной оттепели: ", number)
}
func main() {
	arr1 := []int{-20, 30, -40, 50, 10, -10,12,12,12}
	arr2 := []int{10, 20, 30, 1, -10, 1, 2, 3}
	arr3 := []int{-10, 0, -10, 0, -10}
	printB(arr1)
	printB(arr2)
	printB(arr3)
}