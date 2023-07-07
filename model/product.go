package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
