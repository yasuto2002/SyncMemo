package model

import (
	"context"
	"fmt"
	"log"

	"syncmemo/entity"
	"syncmemo/repository/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeBords(ctx context.Context, db *mongo.Database, item request.Make) string {
	podcastsCollection := db.Collection("boards")
	board := entity.Board{
		NAME:     item.Name,
		MAIL:     "fujiya-1-1@gmail.com",
		PASSWORD: item.Password,
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, board)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
	return insertResult.InsertedID.(primitive.ObjectID).Hex()
}

func GetBoardList(ctx context.Context, db *mongo.Database) []bson.M {
	podcastsCollection := db.Collection("boards")
	cursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var podcasts []bson.M
	if err = cursor.All(ctx, &podcasts); err != nil {
		log.Fatal(err)
	}
	return podcasts
}
