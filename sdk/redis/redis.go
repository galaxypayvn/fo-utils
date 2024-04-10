package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string
	Password string
	DB       int
}

type RedisRepo struct {
	RDB *redis.Client
}

func NewRedisRepo(cfg Config) IRedisRepo {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisRepo{
		RDB: client,
	}
}

type IRedisRepo interface {
	GetRepo() *redis.Client
	SetNX(ctx context.Context, key, value string, exp time.Duration) (bool, error)
	Del(ctx context.Context, key string) error
	SetKey(ctx context.Context, key string, value interface{}, expire time.Duration) (err error)
	GetKey(ctx context.Context, key string) (value string, err error)
	SetHash(ctx context.Context, key string, value interface{}, expire time.Duration) (err error)
	GetHash(ctx context.Context, key string) (res map[string]string, err error)
	CheckExist(ctx context.Context, key string) (res bool, err error)
	GetHashByKey(ctx context.Context, key string, field string) (res string, err error)
}

func (r *RedisRepo) GetRepo() *redis.Client {
	return r.RDB
}

func (r *RedisRepo) SetNX(ctx context.Context, key, value string, exp time.Duration) (bool, error) {
	res := r.RDB.SetNX(ctx, key, value, exp)
	result, err := res.Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *RedisRepo) Del(ctx context.Context, key string) error {
	err := r.RDB.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) CheckExist(ctx context.Context, key string) (res bool, err error) {
	exist, err := r.RDB.Exists(ctx, key).Result()
	if err != nil {
		return res, err
	}
	if exist == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *RedisRepo) SetKey(ctx context.Context, key string, value interface{}, expire time.Duration) (err error) {
	err = r.RDB.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return err
}

func (r *RedisRepo) GetKey(ctx context.Context, key string) (value string, err error) {
	value, err = r.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, err
}

func (r *RedisRepo) SetHash(ctx context.Context, key string, value interface{}, expire time.Duration) (err error) {
	err = r.RDB.HSet(ctx, key, value).Err()
	if err != nil {
		return err
	}
	err = r.RDB.Expire(ctx, key, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetHash(ctx context.Context, key string) (res map[string]string, err error) {
	res, err = r.RDB.HGetAll(ctx, key).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *RedisRepo) GetHashByKey(ctx context.Context, key string, field string) (res string, err error) {
	res, err = r.RDB.HGet(ctx, key, field).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}
