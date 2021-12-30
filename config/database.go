package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es la funcion que realiza la conexion  al db */
var MongoCN = ConectDB()
var clientOptions = options.Client().ApplyURI("")

/* ConectDB() es la funcion que realiza la conexion  al db */
func ConectDB() *mongo.Client {

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

	log.Println("Conexion success to DB!!")

	return client
}

/* CheckConnection() es la funcion que realiza un ping a la db */
func CheckConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
