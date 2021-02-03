package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var mongoDB *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	username, password, host, port := "root", "root", "115.159.151.130","27017"
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	opt := options.Client().ApplyURI(url)
	opt.SetMaxPoolSize(50)
	mongoClient, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal("Mongodb connection failed" + err.Error())
	}

	if err = mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("Mongodb connection failed" + err.Error())
	}

	log.Printf("Mongodb connection successed")
	mongoDB = mongoClient.Database("project")
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}