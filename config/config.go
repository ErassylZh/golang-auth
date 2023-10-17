package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Service  *Service
		Database *Database
	}

	Service struct {
		Port string `envconfig:"PORT" default:"8000"`
	}

	Database struct {
		DBUser     string `envconfig:"DB_USER"`
		DBPassword string `envconfig:"DB_PASSWORD"`
		DBName     string `envconfig:"DB_NAME"`
		DBHost     string `envconfig:"DB_HOST"`
		DBPort     string `envconfig:"DB_PORT"`
		SSLMODE    string `envconfig:"SSL_MODE"`
	}
)

var (
	once   sync.Once
	config *Config
)

func GetConfig(envfiles ...string) (*Config, error) {
	var err error
	once.Do(func() {
		_ = godotenv.Load(envfiles...)

		var c Config
		err = envconfig.Process("", &c)
		if err != nil {
			a
			err = fmt.Errorf("error parse config from env variables: %w\n", err)
			return
		}

		config = &c
	})

	return config, err
}
