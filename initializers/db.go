package initializers

import (
	"fmt"
	"github.com/wpcodevo/golang-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	DB.AutoMigrate(&models.Note{})
	log.Println("🚀 Connected Successfully to the Database")
}
