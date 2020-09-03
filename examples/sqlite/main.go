package main

import (
	"gopkg.in/logex.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

func main() {
	// Create DSN
	/*
		dsnExample := "test.db"

		NOTES:
			- You can also use "file::memory:?cache=shared" instead of a path to a file.
			  This will tell SQLite to use a temporary database in system memory.
			  See https://www.sqlite.org/inmemorydb.html
	*/
	dsn := "products.db"
	//dsn := "file::memory:?cache=shared"

	// Open connection
	logex.Infof("Open connection to dsn %s", dsn)
	db, openErr := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{},
	)
	if openErr != nil {
		panic("open failed")
	}

	// Migrate the schema
	logex.Info("Migrate schema")
	migrateErr := db.AutoMigrate(
		&Product{},
		&Producer{},
	)
	if migrateErr != nil {
		panic("automigrate failed")
	}

	// Create
	prodCode := "D42"
	newProd := &Product{
		Code:  prodCode,
		Price: 100,
		Producer: &Producer{
			Name: "galaxy inc",
		},
	}
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
	//db.Delete(&readProd, prodId)
	//logex.Infof("Deleted product with ID %d: %+v", prodId, readProd)
}
