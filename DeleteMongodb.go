// delete value stored in a mongoDB

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gopal476:gopal476@cluster0.jmfv2vw.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	startdb := client.Database("start")
	pdcollection := startdb.Collection("podcasts")
	//deleting the specific data
	result, err := pdcollection.DeleteOne(ctx, bson.M{"author": "Gopal Karki1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The specified database value is deleted ", result.DeletedCount)
}
