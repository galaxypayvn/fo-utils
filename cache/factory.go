package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// CacheFactoryConfig map cấu hình cho từng loại cache
type CacheFactoryConfig map[string]CacheConfig

// CacheFactory quản lý các containers cache và sync containers
type CacheFactory struct {
	configCache      CacheFactoryConfig
	mapContainer     map[string]ICacheVisitor
	mapSyncContainer map[string]ICacheSync
	redisClient      *redis.Client
	disableCache     bool
	enableSync       bool
}

// NewCacheFactory khởi tạo CacheFactory mới
func NewCacheFactory(redisClient *redis.Client, prefixKey string, configCache CacheFactoryConfig) (*CacheFactory, error) {
	cacheFactory := &CacheFactory{
		redisClient:      redisClient,
		mapContainer:     make(map[string]ICacheVisitor),
		mapSyncContainer: make(map[string]ICacheSync),
		configCache:      configCache,
	}

	// Đăng ký các containers
	if err := cacheFactory.registerAllCacheContainers(prefixKey); err != nil {
		return nil, err
	}

	return cacheFactory, nil
}

// Đăng ký toàn bộ cache containers từ config
func (c *CacheFactory) registerAllCacheContainers(prefixKey string) error {
	for namespace, cacheConfig := range c.configCache {
		if err := c.RegisterCacheContainer(prefixKey, namespace, cacheConfig); err != nil {
			return err
		}
	}
	return nil
}

// RegisterCacheContainer đăng ký một container cache
func (c *CacheFactory) RegisterCacheContainer(prefixKey, namespace string, cacheConfig CacheConfig) error {
	// Khởi tạo cache visitor
	cacheObj, err := NewVisitor(c.redisClient, prefixKey, namespace, cacheConfig)
	if err != nil {
		return err
	}
	c.mapContainer[namespace] = cacheObj

	// Nếu sync được bật, khởi tạo cache sync
	if c.enableSync {
		cacheSyncObj, err := NewSync(c.redisClient, prefixKey, namespace, cacheConfig)
		if err != nil {
			return err
		}
		c.mapSyncContainer[namespace] = cacheSyncObj
	}

	return nil
}

// GetContainer lấy cache visitor cho một namespace cụ thể
func (c *CacheFactory) GetContainer(namespace string) ICacheVisitor {
	if c.disableCache {
		return &emptyVisitor{}
	}
	if cacheContainer, ok := c.mapContainer[namespace]; ok {
		return cacheContainer
	}
	return &emptyVisitor{}
}

// GetSyncContainer lấy sync container cho một namespace cụ thể
func (c *CacheFactory) GetSyncContainer(namespace string) ICacheSync {
	if c.disableCache {
		return &emptySync{}
	}
	if cacheContainer, ok := c.mapSyncContainer[namespace]; ok {
		return cacheContainer
	}
	return &emptySync{}
}

// DisableCache vô hiệu hoá cache
func (c *CacheFactory) DisableCache() ICacheFactory {
	c.disableCache = true
	return c
}

// EnableSync bật tính năng sync
func (c *CacheFactory) EnableSync() ICacheFactory {
	c.enableSync = true
	return c
}

// Flush làm sạch tất cả các containers cache
func (c *CacheFactory) Flush(ctx context.Context) error {
	for _, container := range c.mapContainer {
		if err := container.Reset(ctx); err != nil {
			return err
		}
	}
	return nil
}
