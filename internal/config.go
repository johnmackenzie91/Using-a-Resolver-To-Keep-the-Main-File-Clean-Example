package internal

import "no_vcs/me/resolver/internal/db"

// Config ties together all other application configuration types.
type Config struct {
	DB db.Config
}

func NewConfiguration() *Config {
	return &Config{
		DB: db.GetConfig(),
	}
}

