package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config структура для хранения конфигурации
type Config struct {
	Address        string        `mapstructure:"address"`
	ReportInterval time.Duration `mapstructure:"report_interval"`
	PollInterval   time.Duration `mapstructure:"poll_interval"`
}

// AgentLoadConfig загружает конфигурацию из .env, переменных окружения и задает значения по умолчанию
func AgentLoadConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("ADDRESS", "localhost:8080")
	viper.SetDefault("REPORT_INTERVAL", 10*time.Second)
	viper.SetDefault("POLL_INTERVAL", 2*time.Second)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Файл конфигурации не найден: %v.", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
