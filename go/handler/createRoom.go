package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateRoom struct {
	DB *mongo.Database
}
