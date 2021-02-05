package util

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"userServer/model"
)

func GetAllRoles() []model.RoleDB {
	var allRoles []model.RoleDB
	cursor := FindAll("Role", bson.D{}, nil)
	err := cursor.All(context.TODO(), &allRoles)
	if err != nil {
		log.Println(err.Error())
	}
	err = cursor.Close(context.TODO())
	if err != nil {
		log.Println(err.Error())
	}
	return allRoles
}

func UpdateRole(role model.RoleDB) bool{
	filter := bson.D{{"key", role.Key}}
	update := bson.D{{"$set", bson.D{{"name", role.Name},{"introduction", role.Introduction}}}}
	UpdateOne("Role", filter, update)
	return true
}
