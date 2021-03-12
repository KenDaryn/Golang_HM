package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d", &n)
	arr := []int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr=append(arr,a)
	}
	var b int
	fmt.Println("Введите число: ")
	fmt.Scanf("%d",&b)
	a:=[]int{}

	for i := 0; i < n; i++ {
		if b!=arr[i]{
			a=append(a,arr[i])

		}
	}
	fmt.Println(a)
}
