package handler

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CreateRoom struct {
	DB  *mongo.Database
	CTX context.Context
}
