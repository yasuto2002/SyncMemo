package model

import (
	"context"
	"fmt"
	"math/rand"
	"syncmemo/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Reg(ctx context.Context, db *mongo.Database, user *entity.User) error {
	usersCollection := db.Collection("provisional")
	if _, err := usersCollection.InsertOne(ctx, user); err != nil {
		return fmt.Errorf("provisional insert : %w", err)
	}
	return nil
}

func randNum(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result
}

func EmailCore(ctx context.Context, db *mongo.Database, mail string) error {
	usersCollection := db.Collection("provisional")
	filter := &bson.M{"mail": mail}
	var user entity.User
	result := usersCollection.FindOne(ctx, filter)
	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			// 空の時
			return err
		}
		return err
	}
	return nil
}
