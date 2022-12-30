package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(info string, ctx context.Context) (*mongo.Database, *mongo.Client) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(info))

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	database := client.Database("test")
	return database, client
}
