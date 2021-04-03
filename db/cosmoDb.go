package cosmoDb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	var stringUrl string = "mongodb://antares-padaria:dWThknIXmCHIfydERRpRi26T6eXc6fb2UpBiZUnWCmHvCxJvGjNfWwCCg3eiLkmpB2OKLKw6iZNefsijQPqnyg==@antares-padaria.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@antares-padaria@"
	clientOptions := options.Client().ApplyURI(stringUrl) // Connect to //MongoDB
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