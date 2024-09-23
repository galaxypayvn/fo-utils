package cache

import (
	"context"
)

// Struct chung cho các hành vi no-op
type emptyCacheBehavior struct{}

// Các phương thức no-op cho emptyCacheBehavior
func (e emptyCacheBehavior) Get(ctx context.Context, key string) (interface{}, error) {
	return nil, ErrCacheNotFound
}

func (e emptyCacheBehavior) Set(ctx context.Context, key string, entry interface{}) error {
	return nil
}

func (e emptyCacheBehavior) SetString(ctx context.Context, key string, entry string) error {
	return nil
}

func (e emptyCacheBehavior) Delete(ctx context.Context, key string) error {
	return nil
}

func (e emptyCacheBehavior) Reset(ctx context.Context) error {
	return nil
}

func (e emptyCacheBehavior) SetBulk(ctx context.Context, entries map[string]interface{}) error {
	return nil
}

func (e emptyCacheBehavior) GetList(ctx context.Context, keys []string) ([]interface{}, []string, error) {
	return nil, keys, nil
}

func (e emptyCacheBehavior) GetMap(ctx context.Context, keys []string) (map[string]interface{}, []string, error) {
	return nil, keys, nil
}

func (e emptyCacheBehavior) Push(ctx context.Context, data ...string) error {
	return nil
}

func (e emptyCacheBehavior) Pop(ctx context.Context, count int) ([]string, error) {
	return nil, nil
}

// cacheImplEmpty kế thừa từ emptyCacheBehavior
type cacheImplEmpty struct {
	emptyCacheBehavior
}

var _ ICacheVisitor = (*CacheVisitor)(nil)

// emptyVisitor cũng kế thừa từ emptyCacheBehavior nhưng override Delete
type emptyVisitor struct {
	emptyCacheBehavior
}

// Override hàm Delete trong emptyVisitor để phù hợp với ICacheVisitor interface
func (e emptyVisitor) Delete(ctx context.Context, key string, opt ...bool) error {
	// No-op method for handling Delete with optional arguments
	return nil
}

func (e emptyVisitor) Model(out interface{}) ICacheVisitor {
	return e
}
