package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func NewMongoDatabase(databaseURI string, connTimeout time.Duration) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	cs, err := connstring.Parse(databaseURI)
	if err != nil {
		log.Fatal(err)
	}

	clientOpts := options.Client().ApplyURI(cs.String())

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(cs.Database)
}
