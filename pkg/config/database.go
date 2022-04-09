package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx context.Context

func ConfigDb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/covid_dataset?authSource=admin")
	Ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetColleciton(conn *mongo.Client, database string, collection string) *mongo.Collection {
	return conn.Database(database).Collection(collection)
}
