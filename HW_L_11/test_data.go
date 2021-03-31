package main
//5) test.go
//создание данных клиента и запись их в массив
//массив преобразовать в json и записать в файл(clients.json)
import (
	"Daryn_GO/HW_L_11/models"
	"encoding/json"
	"fmt"
	"os"
)
func main() {
	var cl1, cl2, cl3, cl4 models.Clients
	cl1 = models.Clients{
		Id:       001,
		FullName: "Kenzhebekov Daryn",
		Balance:  999999,
	}
	cl2 = models.Clients{
		Id:       002,
		FullName: "Tasbolatova Saule",
		Balance:  10000,
	}
	cl3 = models.Clients{
		Id:       003,
		FullName: "Erlik Batyr",
		Balance:  88000,
	}
	cl4=models.Clients{
		Id:       004,
		FullName: "Erlik Medina",
		Balance:  65211,
	}
	clients := []models.Clients{cl1, cl2, cl3,cl4}
	fmt.Println(clients)
	clientsJson, err := json.Marshal(clients)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("clients.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(clientsJson)
	if err != nil {
		panic(err)
	}
}