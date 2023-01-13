package sercice

import (
	"context"
	"fmt"
	"syncmemo/entity"
	"syncmemo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login(ctx context.Context, db *mongo.Database, mail string, pw string) (entity.User, error) {
	user, err := model.GetUser(ctx, db, mail)
	if err != nil {
		return entity.User{}, err
	}
	if err := user.ComparePassword(pw); err != nil {
		return entity.User{}, fmt.Errorf("wrong password: %w", err)
	}
	return user, err
}
