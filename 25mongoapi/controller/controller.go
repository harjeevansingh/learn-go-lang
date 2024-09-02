package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/harjeevansingh/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://harjeevansingh01:qwerty1234@go-learn.fzciq.mongodb.net/?retryWrites=true&w=majority&appName=go-learn"
const dbName = "netflix"
const collectionName = "movies"


// Most Important
var collection *mongo.Collection

// Connect to MongoDB
func init(){
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection instance created")
}

// MongoDB Helpers - file


// Insert one record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with ID:", inserted.InsertedID)
}

// Update one record
func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Updated count:", result.ModifiedCount)
}