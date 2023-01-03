package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*Client is Struct Of MongoConnector*/
type Client struct {
	client *mongo.Client
}

var mongoClient = Client{
	client: nil,
}

/*Init MongoDB Conncetion */
func Init() {
	if mongoClient.client != nil {
		fmt.Println("mongoClient is nil")
	}
	fmt.Println("mongoClient", mongoClient)

	mongoUser := os.Getenv("MONGO_USER")
	if mongoUser == "" {
		log.Fatalln("MONGO_USER unset")
	}

	mongoPassword := os.Getenv("MONGO_PASSWORD")
	if mongoPassword == "" {
		log.Fatalln("MONGO_PASSWORD unset")
	}

	mongoEndpoint := os.Getenv("MONGO_ENDPOINT")
	if mongoEndpoint == "" {
		log.Fatalln("MONGO_ENDPOINT unset")
	}

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", mongoUser, mongoPassword, mongoEndpoint)

	c, errMgConnect := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if errMgConnect != nil {
		log.Fatalln(errMgConnect.Error())
	}
	mongoClient.client = c

	fmt.Println("mongoClient", mongoClient)

}

// GetClient return client mongodb
func GetClient() (*Client, error) {
	fmt.Println("mongoClient", mongoClient)
	if mongoClient.client != nil {
		// Ping the primary
		if err := mongoClient.client.Ping(context.TODO(), readpref.Primary()); err != nil {
			return nil, errors.New("Lost Connection")
		}
		return &mongoClient, nil
	}
	return nil, errors.New("Not Connect")
}

// Deconnect mongodb connenction
func (c *Client) Deconnect() error {
	if c.client != nil {
		if err := c.client.Disconnect(context.TODO()); err != nil {
			return errors.New("Deconnect mongo db error")
		}
		fmt.Println("Decconnect mongo db success")
		return nil
	}
	return errors.New("Not Connect")
}
