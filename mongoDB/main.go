package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name string
	Age  int
}

type Test struct {
	NameDes string
	Age     int
	Data    Data
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("AnotherTest").Collection("test")
	res := Test{}
	err = collection.FindOne(context.TODO(), bson.D{{"NameDes", "shyhao"}}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
