package store

import (
	"context"
	"syncmemo/entity"
	"time"

	"github.com/go-redis/redis/v9"
)

func NewAuthKvs(ctx context.Context, addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func TokenSet(ctx context.Context, rdb *redis.Client, user entity.User, token string) error {
	data := &entity.Casual{
		User:  user,
		Token: token,
	}
	err := rdb.Set(ctx, user.MAIL, data, 30*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func TokenGet(ctx context.Context, rdb *redis.Client, mail string) (*entity.Casual, error) {
	result := &entity.Casual{}
	if err := rdb.Get(ctx, mail).Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}

func TokenDelete(ctx context.Context, rdb *redis.Client, key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
