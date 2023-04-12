package config

import (
	"chap2-project2/models"
	"fmt"

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

	db.Debug().AutoMigrate(models.Book{})

	return db
}
