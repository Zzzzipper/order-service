package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		MaxOpenConns int    `mapstructure:"DB_MAX_OPEN_CONNS"`
		MaxIdleConns int    `mapstructure:"DB_MAX_IDLE_CONNS"`
		Host         string `mapstructure:"DB_HOST"`
		Port         int    `mapstructure:"DB_PORT"`
		User         string `mapstructure:"DB_USER"`
		Password     string `mapstructure:"DB_PASSWORD"`
		DbName       string `mapstructure:"DB_NAME"`
		SSLMode      string `mapstructure:"DB_SSLMODE"`
		LogLevel     string `mapstructure:"LOG_LEVEL"`
		GrpcPort     string `mapstructure:"PORT"`
		LogBufUrl    string `mapstructure:"LOGBUFF_URL"`
		LogSource    string `mapstructure:"LOG_SOURCE"`
	}
)

func NewConfig() (*Config, error) {

	cfg := Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
