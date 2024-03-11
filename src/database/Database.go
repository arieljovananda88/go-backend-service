package database

import (
	"fmt"
	"os"
	"sync"

	"go-backend-service/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func initialize() {
	var err error

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)
	print(dsn)

	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	dbInstance.AutoMigrate(&models.Users{})
	dbInstance.AutoMigrate(&models.SystemLogs{})
	dbInstance.AutoMigrate(&models.BimbinganLogs{})
}

func GetInstance() *gorm.DB {
	lock := &sync.Mutex{}

	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if dbInstance == nil {
			initialize()
		}
	}

	return dbInstance
}
