package main

import (
	"fmt"
)
//1) создать структуру клиента банка
//Client
//-Id:int
//-FullName:string
//-Balance:int

type Client struct {
	id int
	FullName string
	Balance int
}

func (cl *Client) IncreaseBalance(){
	cl.Balance=cl.Balance+100
	fmt.Println(cl.id,cl.FullName,cl.Balance)
}
func (cl *Client) DecreaseBalance(){
	cl.Balance=cl.Balance-100
	fmt.Println(cl.id,cl.FullName,cl.Balance)
}
//2) создать методы для структуры
//Client
//-IncreaseBalance(value int) -> увеличивать баланс на value
//-DecreaseBalance(value int) -> уменьшить баланс на value
//-PrintAll() -> выводить данные о клиенте
 func (cl *Client) All(){
	fmt.Println(cl.id, cl.FullName, cl.Balance)

 }


//3) упаковать структуру в пакет models
//надо подумать что тут написать

//4) создать 4 файла increase.go, read.go, decrease.go, test_data.go

//func main() {
//	file, err := os.Create("increase.go")
//	if err != nil {
//		panic(err)
//	}
//	//increase:= IncreaseBalance()
//	n, err := file.WriteString(increase)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(n)
//}

//1) создали структуру +
//2) прописали методы+
//3) создали данные и поместили их в массив+
//4) данные из массива перевели в json format (marhal)+
//5) данные в json формате записали в файл json+
//6) считали данные из json файла и перевели их обратно в массив (unmarhal)+


//5) test.go+
//создание данных клиента и запись их в массив
//массив преобразовать в json и записать в файл(clients.json)

//6) read.go+
//считывать данные из clients.json и выводить через цикл

//7) decrease.go
//считывать данные из client.json и уменьшать balance каждого клиента на 100
//и затем заново все записать в файл client.json

//8) increase.go

//Client1
//id:1
//FullName:"asdad",
//Balance:1000,

//go run test.go
//go run read.go -> увидеть данные о клиенте balance:1000
//go run decrease.go
//go run read.go ->  balance 900