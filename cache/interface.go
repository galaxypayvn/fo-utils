package cache

import (
	"context"
	"errors"
)

// Error messages
var ErrCacheNotFound = errors.New("cache not found")

// Interfaces
type ICacheFactory interface {
	GetContainer(namespace string) ICacheVisitor
	GetSyncContainer(namespace string) ICacheSync
	DisableCache() ICacheFactory
	EnableSync() ICacheFactory
	Flush(ctx context.Context) error
}

type ICacheVisitor interface {
	Get(ctx context.Context, key string) (interface{}, error)
	GetList(ctx context.Context, keys []string) ([]interface{}, []string, error)
	GetMap(ctx context.Context, keys []string) (map[string]interface{}, []string, error)
	Set(ctx context.Context, key string, entry interface{}) error
	SetString(ctx context.Context, key string, entry string) error
	SetBulk(ctx context.Context, entries map[string]interface{}) error
	Delete(ctx context.Context, key string, opt ...bool) error
	Reset(ctx context.Context) error
	Model(out interface{}) ICacheVisitor
}

type ICacheRepository interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, entry interface{}) error
	SetString(ctx context.Context, key string, entry string) error
	SetBulk(ctx context.Context, entries map[string]interface{}) error
	Delete(ctx context.Context, key string) error
	Reset(ctx context.Context) error
	GetList(ctx context.Context, keys []string) ([]interface{}, []string, error)
	GetMap(ctx context.Context, keys []string) (map[string]interface{}, []string, error)
}

type ICacheSync interface {
	GetListKeys(ctx context.Context, limit int) ([]string, error)
	SetVersion(ctx context.Context, key ...string) error
}

type ISyncRepository interface {
	Push(ctx context.Context, data ...string) error
	Pop(ctx context.Context, len int) ([]string, error)
}

// CallBack
type CallbackMem func(context.Context, interface{}, ICacheRepository)
