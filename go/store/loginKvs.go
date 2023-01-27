package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ErrNotFound = errors.New("not found")
)

func NewLoginKvs(ctx context.Context, addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	return rdb
}
func Save(ctx context.Context, rdb *redis.Client, key string, mail string) error {
	err := rdb.Set(ctx, key, mail, 30*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func Load(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	mail, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get by %q: %w", key, ErrNotFound)
	}
	return mail, nil
}

// ログイントークン削除
func LoginDelete(ctx context.Context, rdb *redis.Client, key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
