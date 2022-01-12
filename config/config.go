package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	CatalogServiceHost string
	CatalogServicePort int
	OrderServiceHost   string
	OrderServicePort   int

	CtxTimeout int

	LogLevel string
	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("EVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "127.0.0.1"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefault("CATALOG_SERVICE_PORT", 9000))
	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "127.0.0.1"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 9001))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
