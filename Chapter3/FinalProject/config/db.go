package config

import (
	"fmt"

	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "root"
		dbname   = "db-go-sql"
	)

	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(models.SocialMedia{}, models.User{}, models.Photo{}, models.Comment{})

	return db
}

//export PATH=$PATH:$(go env GOPATH)/bin
