package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code     string
	Price    uint
	Producer *Producer
}

type Producer struct {
	gorm.Model
	ProductID uint
	Name      string
}
