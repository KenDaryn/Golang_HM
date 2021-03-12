package main

import (
	"fmt"
	"math"
)

//import "fmt"
//
//func main(){
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := []int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr=append(arr,a)
//	}
//	var b int
//	fmt.Println("Введите число: ")
//	fmt.Scanf("%d",&b)
//	a:=[]int{}
//
//	for i := 0; i < n; i++ {
//		if b!=arr[i]{
//			a=append(a,arr[i])
//
//		}
//	}
//	fmt.Println(a)
//}

//func printHello(){
//	fmt.Println("Hello world")
//}
//
//func getSumofNumber(){
//	a:=3
//	b:=4
//	c:=a+b
//	fmt.Println(c)
//}
//
//func main(){
////	Функция и процедура
////	Процедура это неки блок когда который выполняется к нему по имени процедуры
//	printHello()
//	getSumofNumber()
//}

//аргументы - входные параметры которые необходимы для работы процедуры или функции

//func printDate(word string){
//	fmt.Println(word)
//}
//
//func getSum(a int, b int){
//	c:=a+b
//	fmt.Println(c)
//}
//func main(){
//	printDate("Hello Daryn")
//	getSum(4,5)
//}

//func printarr(a []int){
//	for i:=0;i<len(a);i++{
//		fmt.Println(a[i])
//	}
//}
//
//func main(){
//	arr1:=[]int{1,2,3,4,5}
//	printarr(arr1)
//	arr2:=[]int{3,4,5,6,7,1}
//	printarr(arr2)
//	arr3:=[]int{55,33,45,4,3}
//	printarr(arr3)
//}

//func printarr(a []int){
//	max:=a[0]
//	min:=a[0]
//	for i:=0;i<len(a);i++{
//		if a[i]>max{
//			max=a[i]
//		} else if a[i]<min{
//			min=a[i]
//		}
//	}
//	fmt.Println(max,min)
//}

//func main(){
//	arr1:=[]int{-1,-2,-3,-4,-5}
//	printarr(arr1)
//	arr2:=[]int{-3,-4,-5,-6,-7,-1}
//	printarr(arr2)
//	arr3:=[]int{-55,-33,-45,-4,-3}
//	printarr(arr3)
//}

//func название функции(параметры....) тип элемента который будет возвращен{
//блок кода
//return то что мы хотим вернуть
//}

//func getMax(a []int) (int, int) {
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
//func getAVG(a[]int) int{
//	sumi:=0
//	avg:=0
//	for i:=0;i<len(a);i++{
//		sumi+=a[i]
//	}
//	avg=sumi/len(a)
//	return avg
//}
//
//func printInfo(a[]int){
//	maxi,mini:=getMax(a)
//	avg:=getAVG(a)
//	fmt.Println(maxi)
//	fmt.Println(mini)
//	fmt.Println(avg)
//}
//
//func main() {
//	arr1 := []int{1, 2, 3, 4}
//	arr2 := []int{4, 5, 6, 2, 323, 23, 4, 232}
//	arr3 := []int{1, 23, 4, 123, 23213, 43434, 4545, 12321, 1313213, 343443, 54565}
//	printInfo(arr1)
//	printInfo(arr2)
//	printInfo(arr3)
//}

//func checkEmail(email string) bool {
//	isCorrect := false
//	for i := 0; i < len(email); i++ {
//		if email[i] == '@' {
//			isCorrect = true
//			break
//		}
//	}
//	return isCorrect
//}
//
//func main() {
//	var email string
//	fmt.Scanf("%s", &email)
//	if checkEmail(email) {
//		fmt.Println("ok")
//	} else {
//		fmt.Println("should contain @")
//	}
//}

// Написать функцию и процедуру
//Макисмальное, минимальное, среднее, четное, нечетные,
//каждый элемент массива возвести в степень n
//(создать функцию которая будет проверять есть ли в массиве какой либо эелемент) bool
//1 2 3 4 5
//1 2 9 16 25 36
//одна единная процедура будет вызваться изз main и выдает все эти записи

func proverka(a []int) bool {
	var nn int
	fmt.Scanf("%d", &nn)
	proverka:=false
	for i := 0; i < len(a); i++ {
		if nn == a[i]{
			proverka=true
			break
		}
	}
	return proverka
}

func getStepen(a []int) []int {
	var n int
	fmt.Scanf("%d", &n)
	stepen := []int{}
	for i := 0; i < len(a); i++ {
		k:=math.Pow(float64(a[i]),float64(n))
		stepen = append(stepen, int(k))
	}
	return stepen
}

func getMaxiMini(a []int) (int, int) {
	maxi := a[0]
	mini := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
		}
		if a[i] < mini {
			mini = a[i]
		}
	}
	return maxi, mini
}

func getAVG(a []int) int {
	sumi := 0
	avg := 0
	for i := 0; i < len(a); i++ {
		sumi += a[i]
	}
	avg = sumi / len(a)
	return avg
}

func getChetNechet(a []int) ([]int, []int) {

	chet := []int{}
	nechet := []int{}
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			chet = append(chet, a[i])
		} else {
			nechet = append(nechet, a[i])
		}
	}
	return chet, nechet
}

func printInfo(a []int)
	maxi, mini := getMaxiMini(a)
	avg := getAVG(a)
	chet, nechet := getChetNechet(a)
	stepen:=getStepen(a)
	proverka:=proverka(a)
	fmt.Println("Максимальное число: ", maxi)
	fmt.Println("Минимальное число: ", mini)
	fmt.Println("Среднее число: ", avg)
	fmt.Println("Четное числa: ", chet, "Не четные числв: ", nechet)
	fmt.Println("Степень: ", stepen)
	fmt.Println("Проверка:  ", proverka)
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 2}
	arr3 := []int{1, 23, 4, 5}
	printInfo(arr1)
	printInfo(arr2)
	printInfo(arr3)
}