package config

import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPserver struct {
	Address string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type Config struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HTTPserver   `yaml:"http_server"`
}

func MustLoad() *Config {
	var config_path string

	config_path = os.Getenv("CONFIG_PATH")

	if config_path == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		config_path = *flags

		if config_path == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", config_path)
	}

	var cfg Config

	err := cleanenv.ReadConfig(config_path, &cfg)

	if err != nil {
		log.Fatalf("Cannot read config file : %s", err.Error())
	}

	return &cfg

}
