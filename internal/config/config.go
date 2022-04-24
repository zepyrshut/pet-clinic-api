package config

import (
	"log"
)

// Application properties
type AppStatus struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

// Application configuration
type AppConfig struct {
	Port    string
	Env     string
	Session struct {
		Secret string
	}
	DB struct {
		DSN string
	}
}

// Application wrapper
type Application struct {
	Status       AppStatus
	Config       AppConfig
	InProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}
