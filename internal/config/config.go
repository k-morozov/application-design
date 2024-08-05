package config

import (
	"flag"
)

const (
	DefaultPort     = "8080"
	DefaultLogLevel = "debug"
)

type ServiceConfig struct {
	Port     string
	LogLevel string
	UseHTTPS bool
}

func ParseConfig() ServiceConfig {
	config := ServiceConfig{}

	flag.StringVar(&config.LogLevel, "l", DefaultLogLevel, "")
	flag.StringVar(&config.Port, "p", DefaultPort, "")

	flag.Parse()
	return config
}

func NewServiceConfigForDebug() ServiceConfig {
	config := ServiceConfig{}
	return config
}
