package main

import (
	"Daryn_GO/15/config"
	"Daryn_GO/15/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

//go get go.mongodb.org/mongo-driver/mongo
func main() {
	//connect to database
	//create config
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "Online_shop",
		CollectionName: "Product",
	}
	//config for client
	clientOptions := options.Client().ApplyURI("mongodb://" + cf.Host + ":" + cf.Port)
	//connect to database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//check for connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//connect to database
	db := client.Database(cf.Database)
	//create collection
	err = db.CreateCollection(context.TODO(), cf.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		log.Fatal(err.Error())
	}
	//create connection to collection
	collection := db.Collection(cf.CollectionName)
	//insert first element
	product1 := models.Product{
		Name:        "BMW",
		Price:       45322.12,
		Description: "German car",
	}

	result, err := collection.InsertOne(context.TODO(), product1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	//get all data from db
	products := []models.Product{}
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range products {
		fmt.Println(v.Name, v.Price,v.Description)
	}
}