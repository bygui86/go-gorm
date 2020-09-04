package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code     string
	Price    uint
	Producer *Producer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Producer struct {
	gorm.Model
	ProductID uint
	Name      string
}
