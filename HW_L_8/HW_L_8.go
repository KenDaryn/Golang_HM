package main

import (
	"fmt"
	"math"
)

func proverka(a []int) bool {
	var nn int
	fmt.Scanf("%d", &nn)
		proverka := false
	for i := 0; i < len(a); i++ {
		if nn == a[i] {
			proverka = true
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
		k := math.Pow(float64(a[i]), float64(n))
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

func printInfo(a []int) {
	maxi, mini := getMaxiMini(a)
	avg := getAVG(a)
	chet, nechet := getChetNechet(a)
	stepen := getStepen(a)
	proverka := proverka(a)
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
