package main

//
//import (
//	"fmt"
//	"math"
//)
//
//func proverka(a []int) bool {
//	var nn int
//	fmt.Scanf("%d", &nn)
//		proverka := false
//	for i := 0; i < len(a); i++ {
//		if nn == a[i] {
//			proverka = true
//			break
//		}
//	}
//	return proverka
//}
//
//func getStepen(a []int) []int {
//	var n int
//	fmt.Scanf("%d", &n)
//	stepen := []int{}
//	for i := 0; i < len(a); i++ {
//		k := math.Pow(float64(a[i]), float64(n))
//		stepen = append(stepen, int(k))
//	}
//	return stepen
//}
//
//func getMaxiMini(a []int) (int, int) {
//	maxi := a[0]
//	mini := a[0]
//	for i := 0; i < len(a); i++ {
//		if a[i] > maxi {
//			maxi = a[i]
//		}
//		if a[i] < mini {
//			mini = a[i]
//		}
//	}
//	return maxi, mini
//}
//
//func getAVG(a []int) int {
//	sumi := 0
//	avg := 0
//	for i := 0; i < len(a); i++ {
//		sumi += a[i]
//	}
//	avg = sumi / len(a)
//	return avg
//}
//
//func getChetNechet(a []int) ([]int, []int) {
//
//	chet := []int{}
//	nechet := []int{}
//	for i := 0; i < len(a); i++ {
//		if a[i]%2 == 0 {
//			chet = append(chet, a[i])
//		} else {
//			nechet = append(nechet, a[i])
//		}
//	}
//	return chet, nechet
//}
//
//func printInfo(a []int) {
//	maxi, mini := getMaxiMini(a)
//	avg := getAVG(a)
//	chet, nechet := getChetNechet(a)
//	stepen := getStepen(a)
//	proverka := proverka(a)
//	fmt.Println("Максимальное число: ", maxi)
//	fmt.Println("Минимальное число: ", mini)
//	fmt.Println("Среднее число: ", avg)
//	fmt.Println("Четное числa: ", chet, "Не четные числв: ", nechet)
//	fmt.Println("Степень: ", stepen)
//	fmt.Println("Проверка:  ", proverka)
//}
//
//func main() {
//	arr1 := []int{1, 2, 3, 4}
//	arr2 := []int{4, 5, 6, 2}
//	arr3 := []int{1, 23, 4, 5}
//	printInfo(arr1)
//	printInfo(arr2)
//	printInfo(arr3)
//}

//func main(){
//	arr:=[][]int{
//		{1,2,3},{4,5,6},{7,8,9},
//	}
//	sum:=0
//	for i:=0;i<len(arr);i++{
//		for j:=0;j<len(arr[i]);j++{
//		sum+=arr[i][j]
//		}
//	}
//	println(sum)
//	//fmt.Println(arr[0][0])
//	//fmt.Println(arr[1][0])
//	//fmt.Println(arr[2][1])
//	//fmt.Println(arr[2][2])
//}

//func main(){
//	arr:=[][]int{
//		{1,2,3,5},
//		{4,5,6,0},
//		{7,8,9,5},
//	}
//
//	for i:=0;i<len(arr);i++{
//		sum:=0
//		for j:=0;j<len(arr[i]);j++{
//			sum+=arr[i][j]
//		}
//		println(sum/len(arr[i]))
//	}
//
//}

//func main(){
//	arr:=[]int{1,2,3,5}
//	for i:=0;i<len(arr);i++{
//		fmt.Println(i,arr[i])
//		}
//		for i,v:=range arr{
//			fmt.Println(i,v)
//		}
//
//	}

//func main() {
//	a := [][]int{}
//
//	n:=3
//	for i := 0; i < n; i++ {
//		a[i]=append(a[i],1)
//		for j := 0; j < n; i++ {
//			a[j]=append(a[j],1)
//
//		}
//	}
//	fmt.Println(a)
//}

//	for i:=0;i<len(arr);i++{
//		sum:=0
//		for j:=0;j<len(arr[i]);j++{
//			sum+=arr[i][j]

//func main(){
////	тут мы генерируем числа для рандома
//	rand.Seed(time.Now().UnixNano())
//	//выдаст л.бое число до 120
//	k:=rand.Intn(120)
//	fmt.Println(k)
//
//}

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


//https://acmp.ru/index.asp?main=task&id_task=264
//готов
//func main() {
//		var n int
//		fmt.Scanf("%d", &n)
//		arr := [10]int{}
//		for i := 0; i < n; i++ {
//			var z int
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
