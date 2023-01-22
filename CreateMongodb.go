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
	//establishing Connection with MongoDB cluster

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gopal476:gopal476@cluster0.jmfv2vw.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	/*err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)*/
	//Creating a database name and collection
	startdb := client.Database("start")
	pdcollection := startdb.Collection("podcasts")
	//epcollection := startdb.Collection("episodes")
	//inserting value in mongodb clusters collection
	pdresult, err := pdcollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "JavaScript Developer"},
		{Key: "author", Value: "Anukul Bhattarai"},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pdresult.InsertedID)

}
