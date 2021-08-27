package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection
var ctx = context.TODO()

func create_connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func create_user(user User) *mongo.InsertOneResult {
	client := create_connection()
	collection := client.Database("training_db").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// func set_counter() int64 {
// client := create_connection()
// collection := client.Database("training_db").Collection("users")
// result, _ := collection.CountDocuments(ctx, User{})
// return result
// }
func find_user(username string) (User, error) {
	client := create_connection()
	collection := client.Database("training_db").Collection("users")

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	var result User

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}, err
	}
	return result, nil
}
