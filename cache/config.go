package cache

type ProviderConfig struct {
	Enabled bool
	TTL     int
}

type CacheConfig struct {
	Providers map[string]ProviderConfig
}
