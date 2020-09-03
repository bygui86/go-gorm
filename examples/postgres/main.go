package main

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "supersecret"
	dbName   = "products"
	params   = "sslmode=disable TimeZone=UTC"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Create DSN
	/*
		dsnExample := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

		NOTES:
			- We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default
	*/
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d %s",
		username, password, host, port, params)
	// if the database already exists
	//dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s %s",
	//	username, password, host, port, dbName, params)

	// Open connection
	logex.Infof("Open connection to dsn %s", dsn)
	db, openErr := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if openErr != nil {
		panic("open failed")
	}

	// Create DB
	logex.Info("Create database")
	dbErr := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
	if dbErr != nil {
		if !strings.Contains(dbErr.Error(), fmt.Sprintf("database \"%s\" already exists", dbName)) {
			panic("database creation failed")
		}
	}

	// Initialize schema
	logex.Info("Init schema")
	migrateErr := db.AutoMigrate(&Product{})
	if migrateErr != nil {
		panic("schema initialization failed")
	}

	// Create
	prodCode := "D42"
	newProd := &Product{Code: prodCode, Price: 100}
	logex.Infof("Create product: %+v", newProd)
	db.Create(newProd)

	// Read
	logex.Info("Read product")
	var readProd Product
	prodId := 1
	db.First(&readProd, prodId) // find product with integer primary key
	logex.Infof("Product with ID %d: %+v", prodId, readProd)
	db.First(&readProd, "code = ?", prodCode) // find product with code D42
	logex.Infof("Product with code %s: %+v", prodCode, readProd)

	// Update
	logex.Info("Update product")
	// update product's price to 200
	db.Model(&readProd).Update("Price", 200) // single field
	logex.Infof("Product with updated price: %+v", readProd)
	// update multiple fields
	updProd := Product{Price: 200, Code: "F42"}
	db.Model(&readProd).Updates(updProd) // using non-zero fields struct
	logex.Infof("Updated product (struct): %+v", readProd)
	updProdMap := map[string]interface{}{"Price": 200, "Code": "F42"}
	db.Model(&readProd).Updates(updProdMap) // using map
	logex.Infof("Updated product (map): %+v", readProd)

	// Delete
	db.Delete(&readProd, prodId)
	logex.Infof("Deleted product with ID %d: %+v", prodId, readProd)
}
