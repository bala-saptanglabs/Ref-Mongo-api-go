package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct{
	Client *mongo.Client
	Db *mongo.Database
}

var MI MongoInstance

func ConnectDB(){
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err!=nil{
		log.Fatal(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()

	// reusing err so no need of :=
	err = client.Connect(ctx)

	if err!=nil{
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Database is connected")

	MI = MongoInstance{client,client.Database(os.Getenv("DB_NAME"))}
}