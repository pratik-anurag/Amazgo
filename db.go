package main
import (
"context"
"encoding/json"
"fmt"
"log"
"net/http"
"os"

"github.com/joho/godotenv"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	config := GetConfiguration()
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("amazgo")

	return collection
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		Message: err.Error(),
		Code:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.Code)
	w.Write(message)
}

type Configuration struct {
	Port             string
	ConnectionString string
}

func GetConfiguration() Configuration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("CONNECTION_STRING"),
	}

	return configuration
}


