package configs

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ApiConfig struct {
	Port string
}

type Config struct {
	DatabaseConfig DatabaseConfig
	ApiConfig      ApiConfig
}

func GetConfig() (config Config) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file")
	}

	apiPort := viper.GetString("api.port")

	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbDatabase := viper.GetString("db.database")

	config = Config{
		DatabaseConfig: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Database: dbDatabase,
		},
		ApiConfig: ApiConfig{
			Port: apiPort,
		},
	}
	return
}
