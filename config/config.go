package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbUser string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_NAME"`
	GrpcPort int `mapstructure:"GRPC_PORT"`
}

// Функция, принимающая путь, название и тип конфиг-файла, возвращает структуру этого конфига
func LoadConfig(confPath, confName, confType string) (Config, error) {

	viper.AddConfigPath(confPath)
	viper.SetConfigName(confName)
	viper.SetConfigType(confType)
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read config file\nError: %v\n", err)
	}
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal data from config file\nError: %v\n", err)
	}
	return config, nil
}