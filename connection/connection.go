package connection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://admin:password@localhost:27018/"
const db = "jwtauthgo"

var Client *mongo.Client
var Database *mongo.Database

func init() {

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	Client = client
	Database = client.Database(db)

	fmt.Println("Mongo DB is ready")
}

func CloseDBConnection() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
