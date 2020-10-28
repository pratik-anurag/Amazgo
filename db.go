package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func ConnectDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("go_rest_api").Collection("amazgo")
	return collection
}

func GetError(err error, w http.ResponseWriter) {

	var response = ErrorResponse{
		Message: err.Error(),
		Code:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.Code)
	w.Write(message)
	log.Fatal(err.Error())
}


