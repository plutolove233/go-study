package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Student struct {
	Name string
	Age  int
}

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
	s1 := Student{
		Name: "小红",
		Age:  11,
	}
	//s2 := Student{
	//	Name: "小黄",
	//	Age:  10,
	//}
	db, err := ConnectToMongoDB("mongodb://localhost:27017", "test", time.Second, 2)
	if err != nil {
		println(err)
	}
	collection := db.Collection("student")

	// insert
	insertRes, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertRes.InsertedID)

	// update
	filter := bson.D{
		{"name", "小红"},
	}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateRes.MatchedCount, updateRes.ModifiedCount)

	// delete
	// 删除名字是小黄的那个
	deleteRes, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "小黄"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteRes.DeletedCount)

	// query
	var result Student
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}
