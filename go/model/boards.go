package model

import (
	"context"
	"fmt"

	"syncmemo/clock"
	"syncmemo/entity"
	"syncmemo/repository/request"
	"syncmemo/repository/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeBords(ctx context.Context, db *mongo.Database, item request.Make, clock clock.Clocker) (response.Make, error) {
	podcastsCollection := db.Collection("boards")
	board := entity.Board{
		NAME:      item.Name,
		MAIL:      item.Mail,
		PASSWORD:  item.Password,
		CreatedAt: clock.Now(),
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, board)
	if err != nil {
		return response.Make{ID: ""}, fmt.Errorf("MakeBords")
	}
	return response.Make{ID: insertResult.InsertedID.(primitive.ObjectID).Hex()}, nil
}

func GetBoardList(ctx context.Context, db *mongo.Database, mail string) ([]response.BoardList, error) {
	filter := &bson.M{"mail": mail}
	opts := options.Find().SetSort(bson.D{{"createdat", -1}}).SetLimit(4)
	boardsCollection := db.Collection("boards")
	cursor, err := boardsCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("FindError")
	}
	var boards []response.BoardList
	if err = cursor.All(ctx, &boards); err != nil {
		return nil, fmt.Errorf("FindError")
	}
	return boards, nil
}

func GetBoardListAll(ctx context.Context, db *mongo.Database, mail string) ([]response.BoardList, error) {
	filter := &bson.M{"mail": mail}
	opts := options.Find().SetSort(bson.D{{"createdat", -1}})
	boardsCollection := db.Collection("boards")
	cursor, err := boardsCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("FindError")
	}
	var boards []response.BoardList
	if err = cursor.All(ctx, &boards); err != nil {
		return nil, fmt.Errorf("FindError")
	}
	return boards, nil
}
