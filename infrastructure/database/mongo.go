package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type MongoDatabase struct {
	driver *mongo.Database
}

func (db *MongoDatabase) Driver() *mongo.Database {
	return db.driver
}

func newClient(databaseURI string, connTimeout time.Duration) *mongo.Database {
	log.Printf("connecting to database %s", databaseURI)
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

func (db *MongoDatabase) NewClient(databaseURI string, connTimeout time.Duration) {
	driver := db.driver
	log.Println("disconnecting from database")
	if err := driver.Client().Disconnect(context.Background()); err != nil {
		log.Panicf("error when disconnecting from database: %v\n", err)
	}
	db.driver = newClient(databaseURI, connTimeout)
}

func NewMongoDatabase(databaseURI string, connTimeout time.Duration) *MongoDatabase {
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

	return &MongoDatabase{
		client.Database(cs.Database),
	}
}
