package model

import (
	"time"
)

type GormCommons struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

type Product struct {
	//gorm.Model
	GormCommons

	Code     string    `json:"code"`
	Price    uint      `json:"price"`
	Producer *Producer `json:"producer,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Producer struct {
	//gorm.Model
	GormCommons

	ProductID uint   `json:"productId"`
	Name      string `json:"name"`
}
