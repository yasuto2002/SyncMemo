package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (context.Context, *mongo.Database, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test:test@mongo:27017/test"))

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	database := client.Database("test")
	return ctx, database, func() {
		cancel()
		client.Disconnect(ctx)
	}
}
