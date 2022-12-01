package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongoDB(uri string, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}
	return client.Database(name), nil
}

func main() {
	db, err := ConnectToMongoDB("mongodb://localhost:27017", "test", time.Second, 2)
	if err != nil {
		println(err)
	}
	println(db.Name())
}
