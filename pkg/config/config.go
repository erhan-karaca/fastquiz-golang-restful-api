package config

import (
	"fastquiz-api/pkg/utils"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBUser        string
	DBPass        string
	DBHost        string
	DBPort        string
	DBName        string
	FrontEndToken string
	ChatGptAPIKey string
	ChatGptModel  string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(utils.GetRootPath() + "/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.SetDefault("CHATGPT_MODEL", "gpt-4o-mini")

	AppConfig = &Config{
		DBUser:        viper.GetString("DB_USER"),
		DBPass:        viper.GetString("DB_PASS"),
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetString("DB_PORT"),
		DBName:        viper.GetString("DB_NAME"),
		FrontEndToken: viper.GetString("FRONTEND_TOKEN"),
		ChatGptAPIKey: viper.GetString("CHATGPT_API_KEY"),
		ChatGptModel:  viper.GetString("CHATGPT_MODEL"),
	}
}
