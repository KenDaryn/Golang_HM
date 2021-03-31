//6) read.go
//считывать данные из clients.json и выводить через цикл

package main

import (
	"Daryn_GO/HW_L_11/models"
	"encoding/json"
	"io/ioutil"
)

func main() {
	readdata, err := ioutil.ReadFile("clients.json")
	if err != nil {
		panic(err)
	}
	var clients []models.Clients
	err = json.Unmarshal(readdata, &clients)
	if err != nil {
		panic(err)
	}
	for _, v := range clients {
		v.PrintAll()
	}
}
