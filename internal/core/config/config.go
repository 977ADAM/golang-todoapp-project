package coreconfig

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	TimeZone *time.Location
}

func NewConfig() (*Config, error) {
	tz := os.Getenv("TIME_ZONE")
	if tz == "" {
		tz = "UTC"
	}

	zone, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("load time zone: %w", err)
	}

	return &Config{
		TimeZone: zone,
	}, nil
}

func NewConfigMust() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to create config: %v", err))
	}
	return cfg
}
