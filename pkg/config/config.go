package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Database struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Name       string `yaml:"name"`
	Port       int    `yaml:"port"`
	SSLMode    string `yaml:"sslmode"`
	Migrations bool   `yaml:"migrations"`
}

type Config struct {
	Env      string     `yaml:"env" env-default:"local"`
	GRPC     GRPCConfig `yaml:"grpc"`
	Database Database   `yaml:"database"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic(any("Файл конфигурации по указанному пути отсутствует"))
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(configPath string) *Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(any(fmt.Sprintf("Файл конфигурации не найден: %s", configPath)))
	}

	cfg := new(Config)

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic(any(fmt.Sprintf("Ошибка чтения файла конфигурации: %v", err)))
	}

	return cfg
}

// fetchConfigPath - парсинг пути к конфигурации из флага или переменной окружения
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
