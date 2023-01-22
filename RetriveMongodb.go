// read / retrieve or read data from MongoDB

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

	cursor, err := pdcollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//Exampe1:
	//making a slice and then displaying the value
	/*var podcasts []bson.M
	if err = cursor.All(ctx, &podcasts); err != nil {
		log.Fatal(err)
	}
	for _, podcasts := range podcasts {
		// displays all the record from the mongoDB cluster
		//fmt.Println(podcasts)
		// display specific data only
		fmt.Println(podcasts["author"])
	}*/

	//Example 2:
	defer cursor.Close(ctx)
	//do one by one using next
	for cursor.Next(ctx) {
		var podcasts bson.M
		if err = cursor.Decode(&podcasts); err != nil {
			log.Fatal(err)
		}
		fmt.Println(podcasts)
	}
}
