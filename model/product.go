package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	UserID uint   `json:"userID"`
}
