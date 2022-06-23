package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IP   string `env:"L2_IP" env-default:"127.0.0.1"`
	Port string `env:"L2_PORT" env-default:":8080"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(
		func() {
			cfg = &Config{}

			if err := cleanenv.ReadEnv(cfg); err != nil {
				fmt.Printf("environment is not OK: %s\n", err)
				os.Exit(1)
			}
		},
	)

	return cfg
}
