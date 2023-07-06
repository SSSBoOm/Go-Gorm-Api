package database

import (
	"fmt"

	"github.com/SSSBoOm/go-gorm-api/config"
	"github.com/SSSBoOm/go-gorm-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database instance
var db *gorm.DB

// Connect function
func Connect() error {
	var err error

	dsn := config.Config("DB_PORT")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")

	return nil
}
