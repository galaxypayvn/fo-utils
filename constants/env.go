package constants

const (
	EnvProduction  = "production"
	EnvStaging     = "staging"
	EnvDevelopment = "development"
	EnvLocal       = "local"
)

const (
	AppDebug        = "APP_DEBUG"
	AppEnv          = "APP_ENV"
	AppKey          = "APP_KEY"
	AppName         = "APP_NAME"
	AppVersion      = "APP_VERSION"
	AppHost         = "APP_HOST"
	AppPort         = "APP_PORT"
	AppTimezone     = "APP_TIMEZONE"
	AppCluster      = "APP_CLUSTER"
	AppNamespace    = "APP_NAMESPACE"
	AppRegistryAddr = "APP_REGISTRY_ADDR"
	AppRegistryPwd  = "APP_REGISTRY_PWD"
	AppEndpoint     = "APP_ENDPOINT"

	DBEngine       = "DB_ENGINE"
	DBHost         = "DB_HOST"
	DBPort         = "DB_PORT"
	DBHostRW       = "DB_HOST_RW"
	DBPortRW       = "DB_PORT_RW"
	DBHostRO       = "DB_HOST_RO"
	DBPortRO       = "DB_PORT_RO"
	DBUser         = "DB_USER"
	DBPwd          = "DB_PWD"
	DBName         = "DB_NAME"
	DBSSLMode      = "DB_SSL_MODE"
	DBConnStr      = "DB_CONN_STR"
	DBConnLifetime = "DB_CONN_LIFETIME"
	DBConnMaxIdle  = "DB_CONN_MAX_IDLE"
	DBConnMaxOpen  = "DB_CONN_MAX_OPEN"

	APMEnabled                  = "APM_ENABLED"
	APMProviderDatadogAgentHost = "APM_PROVIDER_DATADOG_AGENT_HOST"
	APMProviderDatadogAgentPort = "APM_PROVIDER_DATADOG_AGENT_PORT"
	APMProviderElasticAgentHost = "APM_PROVIDER_ELASTIC_AGENT_HOST"
	APMProviderElasticAgentPort = "APM_PROVIDER_ELASTIC_AGENT_PORT"

	BrokerAddr   = "BROKER_ADDR"
	BrokerKey    = "BROKER_KEY"
	BrokerSecret = "BROKER_SECRET"

	RedisHost = "REDIS_HOST"
	RedisPort = "REDIS_PORT"
	RedisPwd  = "REDIS_PWD"
)
