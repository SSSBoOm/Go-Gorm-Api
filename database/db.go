package database

import (
	"fmt"

	"github.com/SSSBoOm/go-gorm-api/config"
	"github.com/SSSBoOm/go-gorm-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dsn := config.Config("MYSQL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")

	return nil
}
