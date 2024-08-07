package config

import (
	"flag"
	"time"
)

const (
	DefaultPort          = "8080"
	DefaultLogLevel      = "debug"
	DefaultHandleTimeout = 2 * time.Minute
)

const (
	InMemory StorageType = iota
	LocalFile
	Database
)

type StorageType int

type ServiceConfig struct {
	Port          string
	LogLevel      string
	UseHTTPS      bool
	StorageType   StorageType
	HandleTimeout time.Duration
}

func ParseConfig() ServiceConfig {
	config := ServiceConfig{
		HandleTimeout: DefaultHandleTimeout,
	}

	flag.StringVar(&config.LogLevel, "l", DefaultLogLevel, "")
	flag.StringVar(&config.Port, "p", DefaultPort, "")

	flag.Parse()
	return config
}

func NewServiceConfigForDebug() ServiceConfig {
	config := ServiceConfig{
		LogLevel:      DefaultLogLevel,
		HandleTimeout: DefaultHandleTimeout,
	}
	return config
}

func (c ServiceConfig) IsMemoryStorage() bool {
	return c.StorageType == InMemory
}
