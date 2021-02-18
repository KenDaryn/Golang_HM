package Test

import "fmt"
//%d - число
//%s - строка
//это input как в Python
func main ()  {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	c:=a+b
	fmt.Println(c)
}
