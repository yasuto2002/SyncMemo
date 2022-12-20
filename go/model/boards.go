package model

import (
	"context"
	"fmt"

	"syncmemo/entity"
	"syncmemo/repository/request"

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
