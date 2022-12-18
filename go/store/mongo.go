package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(info string) (context.Context, *mongo.Database, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(info))

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
