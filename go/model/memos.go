package model

import (
	"context"
	"fmt"
	"syncmemo/entity"
	"syncmemo/repository/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Memos struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

func MakeMemos(ctx context.Context, db *mongo.Database, memo request.Memo) (string, error) {
	memosCollection := db.Collection("Memos")
	memos := entity.Memo{
		Text:    memo.Text,
		X:       memo.X,
		Y:       memo.Y,
		BoardId: memo.BoardId,
	}
	insertResult, err := memosCollection.InsertOne(ctx, memos)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetMemos(ctx context.Context, db *mongo.Database, boardId string) ([]entity.Memo, error) {
	memosCollection := db.Collection("Memos")
	filter := &bson.M{"boardid": boardId}
	memo, err := memosCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("FindError")
	}
	var memos []entity.Memo
	if err = memo.All(ctx, &memos); err != nil {
		return nil, fmt.Errorf("FindError")
	}
	return memos, nil
}
