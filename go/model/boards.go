package model

import (
	"context"
	"fmt"

	"syncmemo/entity"
	"syncmemo/repository/request"
	"syncmemo/repository/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeBords(ctx context.Context, db *mongo.Database, item request.Make) (response.Make, error) {
	podcastsCollection := db.Collection("boards")
	board := entity.Board{
		NAME:     item.Name,
		MAIL:     "fujiya-1-1@gmail.com",
		PASSWORD: item.Password,
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, board)
	if err != nil {
		return response.Make{ID: ""}, fmt.Errorf("MakeBords")
	}
	return response.Make{ID: insertResult.InsertedID.(primitive.ObjectID).Hex()}, nil
}

func GetBoardList(ctx context.Context, db *mongo.Database) ([]response.BoardList, error) {
	filter := &bson.D{}
	boardsCollection := db.Collection("boards")
	cursor, err := boardsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("FindError")
	}
	var boards []response.BoardList
	if err = cursor.All(ctx, &boards); err != nil {
		return nil, fmt.Errorf("FindError")
	}
	return boards, nil
}
