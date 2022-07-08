package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * The DB interface is used to implement the database operations.
 */
var MongoConnection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://pcattaneo:Pablo123@cluster0.jeblk.mongodb.net/?retryWrites=true&w=majority")

/**
 * The Mongo connection is used to connect to the MongoDB server.
 */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)	
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Connected to Mongo")
	return client
}

func checkConnection() bool {
	err := client.Ping(context.TODO(), nil)	
	if err != nil {
		return false
	}

	return true
}