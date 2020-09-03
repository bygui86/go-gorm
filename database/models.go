package database

import (
	"github.com/bygui86/go-gorm/model"
	"gorm.io/gorm"
)

type DbInterface interface {
	OpenConnection() error
	InitSchema() error

	GetProducts() ([]*model.Product, error)
	GetProductById(uint) (*model.Product, error)
	CreateProduct(*model.Product) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)
	DeleteProduct(uint) error
}

type DbInterfaceImpl struct {
	cfg *config
	db  *gorm.DB
}

type config struct {
	dbType dbType
	dbName string
}

const (
	sqliteDb   dbType = "sqlite"
	mysqlDb    dbType = "mysql"
	postgresDb dbType = "postgres"
)

type dbType string

func (d dbType) string() string {
	return string(d)
}
