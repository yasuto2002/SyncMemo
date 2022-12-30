package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

func MakePodcast(ctx context.Context, db *mongo.Database) error {
	podcastsCollection := db.Collection("podcasts")
	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}
