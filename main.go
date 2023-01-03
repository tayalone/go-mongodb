package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	MongoClient "github.com/tayalone/go-mongodb/mongo"
)

func main() {
	fmt.Println("Let's Try")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoClient.Init()

	client, errMongoClient := MongoClient.GetClient()

	if errMongoClient != nil {
		log.Fatalln("errMongoClient", errMongoClient.Error())
	}
	client.Deconnect()

	// mongoUser := os.Getenv("MONGO_USER")
	// if mongoUser == "" {
	// 	log.Fatalln("MONGO_USER unset")
	// }

	// mongoPassword := os.Getenv("MONGO_PASSWORD")
	// if mongoPassword == "" {
	// 	log.Fatalln("MONGO_PASSWORD unset")
	// }

	// mongoEndpoint := os.Getenv("MONGO_ENDPOINT")
	// if mongoEndpoint == "" {
	// 	log.Fatalln("MONGO_ENDPOINT unset")
	// }

	// uri := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", mongoUser, mongoPassword, mongoEndpoint)

	// _, errMgConnect := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	// if errMgConnect != nil {
	// 	log.Fatalln(errMgConnect.Error())
	// }

}
