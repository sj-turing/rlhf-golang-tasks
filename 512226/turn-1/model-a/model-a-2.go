package main

import (
	"context"
	"fmt"
	"log"
	"mongo-driver/mongo"
	"mongo-driver/mongo/options"
)

type UserInteraction struct {
	ID        mongo.ObjectID `bson:"_id,omitempty"`
	UserID    string         `bson:"user_id"`
	Page      string         `bson:"page"`
	Timestamp int64          `bson:"timestamp"`
}

func storeInteraction(client *mongo.Client, interaction UserInteraction) error {
	collection := client.Database("ecommerce").Collection("interactions")
	_, err := collection.InsertOne(context.TODO(), interaction)
	return err
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	interaction := UserInteraction{
		UserID:    "123",
		Page:      "product_detail",
		Timestamp: 1609459200,
	}

	if err := storeInteraction(client, interaction); err != nil {
		log.Fatalf("failed to store interaction: %v", err)
	}

	fmt.Println("Interaction stored successfully.")
}
