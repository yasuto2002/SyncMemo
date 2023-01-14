package model

import (
	"context"
	"syncmemo/entity"
	"syncmemo/repository/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Memos struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

func MakeMemos(ctx context.Context, db *mongo.Database, memo request.Memo) (string, error) {
	memosCollection := db.Collection("Memos")
	memos := entity.Memos{
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
