package main

import "fmt"

//https://acmp.ru/index.asp?main=task&id_task=264
//готов
//func main() {
//		var n int
//		fmt.Scanf("%d", &n)
//		arr := [10]int{}
//		for i := 0; i < n; i++ {
//			var z intexit
//			fmt.Scanf("%d", &z)
//			arr[i] = z
//		}
//	max := 0
//	ottepel := 0
//	for i := 0; i < n; i++ {
//		if arr[i] < 0 {
//			for j := 0; j < n; j++ {
//				if arr[j] < 0 {
//					ottepel = 0
//					break
//				}
//			}
//		}
//		if arr[i] >= 0 {
//			ottepel += 1
//			if ottepel > max {
//				max += 1
//			}
//		}
//	}
//	fmt.Println("Продолжительная оттепель = ", max)
//}

//https://acmp.ru/index.asp?main=task&id_task=496
//Готово
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [10]int{}
//	for i := 0; i < n; i++ {
//		var z int
//		fmt.Scanf("%d", &z)
//		arr[i] = z
//	}
//	sumi:=0
//	if n%2==0{
//		for i := 1; i < n;i++ {
//		 	sumi+=arr[i]}
//	} else{
//		for i:=0;i<n;i++{
//			sumi+=arr[i]
//		}
//	}
//	fmt.Println(sumi)
//}

//https://acmp.ru/index.asp?main=task&id_task=131
//готово
//func main() {
//	arr2 := [][]int{
//		{25, 1},
//		{70, 1},
//		{100, 0},
//		{3, 1},
//	}
//	max := arr2[0][0]
//	for i := 0; i < len(arr2); i++ {
//		if arr2[i][0]>max {
//			max = arr2[i][0]
//		}
//	}
//	fmt.Println("Максимальный возраст = ",max)
//}

//Вывести неповторяющиеся элементы массива
//В массиве найти элементы, которые в нем встречаются только один раз, и вывести их на экран.
//То есть найти и вывести уникальные элементы массива.
//готово
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [10]int{}
//	for i := 0; i < n; i++ {
//		var z int
//		fmt.Scanf("%d", &z)
//		arr[i] = z
//	}
//	for i:=0;i<n;i++{
//		f:=true
//		for j:=0;j<n;j++{
//			if arr[i]==arr[j] && i!=j{
//				f=false
//				break
//			}
//		}
//		if f==true{
//			fmt.Println(arr[i])
//		}
//	}
//}

//2) Определить строки матрицы, в которых число 5 встречается 3 и более раз
//Матрицу 10x20 заполнить случайными числами от 0 до 15.
//Вывести на экран саму матрицу и номера строк,
//в которых число 5 встречается три и более раз.
//Готово надо сделать ввиде функции
//1 2 3 4
//5 2 1 2
//5 5 5 2
//3 2 3 4
//
//Ответ:2
//Готово
//func CreateTwoDimensionalArray(rows, columns, randomRange int) [][]int {
//	arr := [][]int{}
//	//code
//	for i := 0; i < rows; i++ {
//		a := []int{}
//		for j := 0; j < columns; j++ {
//			k := rand.Intn(randomRange)
//			a = append(a, k)
//		}
//		arr = append(arr, a)
//	}
//	return arr
//}
//
//func CreateArray(n, randomRange int) []int {
//	rand.Seed(time.Now().UnixNano())
//	arr := []int{}
//	for i := 0; i < n; i++ {
//		k := rand.Intn(randomRange)
//		arr = append(arr, k)
//	}
//	return arr
//}
//
//func main() {
//	arr := CreateTwoDimensionalArray(10, 20, 15)
//	for _, v := range arr {
//		fmt.Println(v)
//	}
//	max := 0
//	ii := 0
//	for i := 0; i < len(arr); i++ {
//		shetchik := 0
//		for j := 0; j < len(arr[i]); j++ {
//			if arr[i][j] == 5 {
//				shetchik += 1
//				if shetchik > max {
//					max = shetchik
//					ii = i
//				}
//			}
//		}
//	}
//	fmt.Println("Повторение в массиве = ", max, "", "Строка = ", ii)
//}

//3) Найти массив с максимальной суммой элементов
//Сгенерировать десять массивов из случайных чисел
//Вывести их и сумму их элементов на экран +
//Найти среди них один с максимальной суммой элементов+
//Указать какой он по счету, повторно вывести этот массив и сумму его элементов++.
//Заполнение массива и подсчет суммы его элементов оформить в виде отдельной функции.

//Готово надо сделать ввиде функции

//1 2 3 4 5 65
//1 2 3 4 5 6
//1 2 3 4 5 6
//123 34 1 2 4

//func main(){
//	arr:=[][]int{
//	{1,2,3,4,5,65},
//	{1000,2,3,4,5,6},
//	{1,2,3,4,5,6},
//	{123,34,1,2,4},
//	}
//	max:=0
//	ii:=0
//	for i:=0;i<len(arr);i++{
//		maxisumm:=0
//		for j:=0;j<len(arr[i]);j++{
//			maxisumm=maxisumm+arr[i][j]
//			if maxisumm>max{
//				max=maxisumm
//				ii=i
//			}
//
//		}
//		fmt.Println("Общая сумма каждого массива = ",maxisumm)
//	}
//	fmt.Println("Максимальная сумма массива = ", max)
//	fmt.Println("Строка максимального массива = ", ii)
//}

func FullBoks(s Book) {
	fmt.Println(s.Name, s.Price)
}

type Book struct {
	Name  string
	Price int
}

func main() {
	b1 := Book{
		Name: "Книга №1",
		Price:  1200,

	}
	b2 := Book{
		Name:  "Книга №2",
		Price: 9000,
	}
	students := []Book{b1, b2}
	for _, v := range students {
		FullBoks(v)
	}

}
