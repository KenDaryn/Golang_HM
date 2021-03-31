package main

import (
	"Daryn_GO/HW_L_11/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	err := clients.WriteString()
	if err != nil {
		panic(err)
	}
}



//7) decrease.go
//считывать данные из client.json и уменьшать balance каждого клиента на 100
//и затем заново все записать в файл client.json