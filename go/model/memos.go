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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateMemo(ctx context.Context, db *mongo.Database, memo request.Memo) {
	memos := db.Collection("Memos")
	id, err := primitive.ObjectIDFromHex(memo.Id)
	if err != nil {
		log.Printf("Error insert Memo: %v", err)
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"text", memo.Text},
		{"boardid", memo.BoardId},
		{"x", memo.X},
		{"y", memo.Y},
	}}}
	opts := options.Update().SetUpsert(true)
	if _, err := memos.UpdateOne(ctx, filter, update, opts); err != nil {
		log.Printf("Error insert Memo: %v", err)
	}
}

// チャンネルからmongoに入れていく
func AddCh(ctx context.Context, db *mongo.Database, ch chan request.Memo) {
	for {
		select {
		case memo := <-ch:
			UpdateMemo(ctx, db, memo)
		default:
			//fmt.Println("No value")
		}
	}
}

func DeleteMemo(ctx context.Context, db *mongo.Database, memoId string) error {
	memoCollection := db.Collection("Memos")
	id, err := primitive.ObjectIDFromHex(memoId)
	filter := &bson.M{"_id": id}
	if err != nil {
		return fmt.Errorf("id chenge error")
	}
	if _, err := memoCollection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("BoardDeleteError")
	}
	return nil
}
