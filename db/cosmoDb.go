package cosmoDb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	uri := "mongodb+srv://adoleta:AmDmi0cv64LPABdR@cluster0.kjmlo.mongodb.net/adoleta?retryWrites=true&w=majority"
	
	// var stringUrl string = "mongodb://adoleta:7wQgdY054YQQ2BmJ9GAu1pb0MoqJtJ0f9JabKIDxCzISKIIX7hhNmTS0a2Y1j7QXqYJYna93lAnLFeUGdXqDAQ==@adoleta.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@adoleta@"
	
	clientOptions := options.Client().ApplyURI(uri) // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client

}
