package main

import "fmt"

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

	func main(){
			var n int
			fmt.Scanf("%d", &n)
			arr := [100] int{}
			for i := 0; i < n; i++ {
				var a int
				fmt.Scanf("%d", &a)
				arr[i] = a
	}
		chet := [10] int{}
		nechet:=[10] int{}
		for i := 0; i < n; i++ {
			if arr[i]%2==0{
				chet[i]=arr[i]
			} else{
				nechet[i]=arr[i]
			}
		}
		fmt.Println("Четные = ", chet)
		fmt.Println("Не четные", nechet)
	}