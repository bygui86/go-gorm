package main

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	username = "root"
	password = "supersecret"
	dbName   = "products"
	params   = "charset=utf8mb4&parseTime=True&loc=Local"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Create DSN
	/*
		refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

		dsnExample := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

		NOTES:
			- To handle time.Time correctly, you need to include parseTime as a parameter.
			  For more parameters see https://github.com/go-sql-driver/mysql#parameters
			- To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4.
			  For a detailed explanation see https://mathiasbynens.be/notes/mysql-utf8mb4
	*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		username, password, host, port, dbName, params)

	// Open connection
	logex.Infof("Open connection to dsn %s", dsn)
	db, openErr := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if openErr != nil {
		panic("open failed")
	}

	// Migrate the schema
	logex.Info("Migrate schema")
	migrateErr := db.AutoMigrate(&Product{})
	if migrateErr != nil {
		panic("automigrate failed")
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
