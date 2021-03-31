package models

import "fmt"

type Clients struct {
	Id int
	FullName string
	Balance int
}

func (cl *Clients) PrintAll(){
	fmt.Println(cl.Id, cl.FullName, cl.Balance)
}

func (cl *Clients) IncreaseBalance(a int){
	cl.Balance=cl.Balance+a
	fmt.Println(cl.Id,cl.FullName,cl.Balance)
}
func (cl *Clients) DecreaseBalance(b int){
	cl.Balance=cl.Balance-b
	fmt.Println(cl.Id,cl.FullName,cl.Balance)
}
