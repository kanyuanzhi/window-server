package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"userServer/driver"
)

var mongoDB = driver.GetMongoDB()

func FindOne(collection string, filter bson.D) *mongo.SingleResult {
	singleResult := mongoDB.Collection(collection).FindOne(context.TODO(), filter)
	if singleResult != nil {
		log.Println(singleResult)
	}
	return singleResult
}

func InsertOne(collection string, document interface{}) (insertResult *mongo.InsertOneResult) {
	insertResult, err := mongoDB.Collection(collection).InsertOne(context.TODO(), document)
	log.Println(insertResult)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func UpdateOne(collection string, filter interface{}, update interface{}) (updateResult *mongo.UpdateResult) {
	updateResult, err := mongoDB.Collection(collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return
}
