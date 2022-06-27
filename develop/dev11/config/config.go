package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config представляет собой конфиг запускаемого http сервера
type Config struct {
	IP   string `env:"L2_IP" env-default:"127.0.0.1"`
	Port string `env:"L2_PORT" env-default:":8080"`
}

var cfg *Config
var once sync.Once

// GetConfig получает конфиг из переменных окружения и записывает значения в структуру Config, возвращая его
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
