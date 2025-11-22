package config

import "time"

type Config struct {
	Addr string
	DB   DBConfig
	// Add more config sections as needed
	// Server ServerConfig
	// Features FeatureConfig
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}
