package main

import (
	"Daryn_GO/15/config"
	"Daryn_GO/15/models"
	"fmt"
	"log"
)

func main() {
	cf1 := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "Online_shop",
		CollectionName: "Product",
	}
	collection, err := models.GetConnection(cf1)
	if err != nil {
		log.Fatal(err)
	}
	products, err := models.GetAllProduct(collection)
	fmt.Println(products)
}