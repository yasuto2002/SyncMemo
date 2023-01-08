package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	Mongo string
	Kvs   string
}

func New() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	cfg.Port = os.Getenv("PORT")
	cfg.Mongo = os.Getenv("MONGO")
	cfg.Kvs = os.Getenv("REDIS")
	return cfg, nil
}
