package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dbConfig := GetConfig().DatabaseConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}
