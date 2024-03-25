package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Config struct {
	Env            string            `yaml:"env" env-default:"local"`
	PostgresSQL    PostgresSQLConfig `yaml:"postgres"`
	GRPC           GRPCConfig        `yaml:"grpc"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"TokenTTL" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type PostgresSQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Priority: flag > env > default.
// Default value is empty string.

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path if empty " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

// DSN returns the Data Source Name.
func (c *PostgresSQLConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.Database)
}
