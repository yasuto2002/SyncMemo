package store

import (
	"context"
	"syncmemo/entity"
	"syncmemo/repository/request"
	"time"

	"github.com/go-redis/redis/v9"
)

func NewMemoKvs(ctx context.Context, addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       2,  // use default DB
	})
	return rdb
}

func MemoSet(ctx context.Context, rdb *redis.Client, memo request.Memo, id string) error {
	data := entity.Memo{
		Text:    memo.Text,
		X:       memo.X,
		Y:       memo.Y,
		BoardId: memo.BoardId,
	}
	err := rdb.Set(ctx, id, data, 30*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
func MemoGet(ctx context.Context, rdb *redis.Client, memo request.Memo, id string) error {
	data := entity.Memo{
		Text:    memo.Text,
		X:       memo.X,
		Y:       memo.Y,
		BoardId: memo.BoardId,
	}
	err := rdb.Set(ctx, id, data, 30*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
