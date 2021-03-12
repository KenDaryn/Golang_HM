package main

//1) вводятся любые числа в массив
//и среди них необходимо найти те числе которые являются четными и положить их индексы
//в другой массив и вывести
//
//4 9 2 3 4 5 1
//0 1 2 3 4 5 6
//
//Ответ:0 2 4

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [5]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	arr2 := [5]int{}
//	z:=0
//	for i := 0; i < n; i++ {
//		if arr[i]%2 != 0 {
//			continue
//
//		}
//		arr2[z] = i
//		z+=1
//		fmt.Println("Индексы",i)
//	}
//	fmt.Println("Массив с индексами = ",arr2)
//}

//2) вводятся числа в массив там могут быть и отрицательные и положительные и вам необходимо
//посчитать сумму элементов после первого отрицательного числа
//
//1 2 3 -1 2 3 4
//Ответ:9

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [7]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	number:=0
//	for i:=0;i<n;i++{
//		if arr[i]<0{
//			break
//		}
//		number+=1
//	}
//	sum:=0
//	for i:=number+1;i<n;i++{
//		sum+=arr[i]
//	}
//	fmt.Println("Итого = ",sum)
//}

//3)
//n=5
//4
//3
//10
//2
//4
//На что нужно поделить:2 ->d
//
//необходимо вывести количество чисел из массива которые делятся на d без остатка
//Ответ:4

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	arr := [5]int{}
//	for i := 0; i < n; i++ {
//		var a int
//		fmt.Scanf("%d", &a)
//		arr[i] = a
//	}
//	var d int
//	fmt.Println("Введите число деление: ")
//	fmt.Scanf("%d", &d)
//	sum := 0
//	for i := 0; i < n; i++ {
//		if arr[i]%d == 0 {
//			fmt.Println(arr[i])
//			sum += 1
//		}
//	}
//	fmt.Println("Итого = ", sum)
//}

//задачи - https://acmp.ru

//func main(){
//	a:=[]int{1,2,3,45}
//	a=append(a,3)
//	fmt.Println(a)
//	fmt.Println(len(a))
//}

//func main(){
//	a:=[]int{1,2,3,45}
//	b:=[]int{}
//	c:=[]int{}
//	//a=append(a,3)
//	//fmt.Println(a)
//	//fmt.Println(len(a))
//	for i:=0;i<len(a);i++{
//		if a[i]%2==0
//			b=append(b,a[i])
//	}
}

//func main() {
//	sl := []string{"daryn", "anel", "kairat", "12345", "123121"}
//	a := sl[0:3]
//	//b:=sl[4:5]
//	fmt.Println(a)
//}
