package model

import (
	"context"
	"fmt"
	"log"

	"syncmemo/entity"
	"syncmemo/repository/request"
	"syncmemo/repository/response"

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

func GetBoardList(ctx context.Context, db *mongo.Database) []response.BoardList {
	filter := &bson.D{}
	boardsCollection := db.Collection("boards")
	cursor, err := boardsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	var boards []response.BoardList
	if err = cursor.All(ctx, &boards); err != nil {
		log.Fatal(err)
	}
	return boards
}
