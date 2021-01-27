package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConn connection object */
var MongoConn = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin-twittor:H7hc6adVSBAPXm3@twitter.yepml.mongodb.net/twitter?retryWrites=true&w=majority")

/*ConnectDB is the mothod to connect with our MongoDB database*/
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
	log.Println("Sucessfully connected!")
	return client
}

/*CheckConnection function that return 1 if the connection is ok or 0 if is failed */
func CheckConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
