package main

import (
	"github.com/bygui86/go-gorm/database"
	"github.com/bygui86/go-gorm/model"
	"gopkg.in/logex.v1"
)

const (
	productCode  = "D42"
	productPrice = 100
	producerName = "galaxy inc"

	newCode  = "F42"
	newPrice = 200
)

var (
	db *database.DbInterfaceImpl
)

func main() {
	logex.Info("go-gorm start")

	db = newDb()

	openConnection()

	initSchema()

	product := createProduct()

	getProductById(product.ID)

	updateProduct(product)

	getProductById(product.ID)

	deleteProduct(product.ID)

	getAllProducts()

	logex.Info("go-gorm completed")
}

func newDb() *database.DbInterfaceImpl {
	logex.Info("Create new database interface")
	db, err := database.NewDbInterface()
	if err != nil {
		logex.Fatal(err)
	}
	return db
}

func openConnection() {
	logex.Info("Open database connection")
	err := db.OpenConnection()
	if err != nil {
		logex.Fatal(err)
	}
}

func initSchema() {
	logex.Info("Initialize database schema")
	err := db.InitSchema()
	if err != nil {
		logex.Fatal(err)
	}
}

func createProduct() *model.Product {
	product, err := db.CreateProduct(&model.Product{
		Code:  productCode,
		Price: productPrice,
		Producer: &model.Producer{
			Name: producerName,
		},
	})
	if err != nil {
		logex.Fatal(err)
	}
	logex.Infof("Product successfully created: %+v", product)
	return product
}

func getProductById(productId uint) {
	foundProduct, err := db.GetProductById(productId)
	if err != nil {
		logex.Fatal(err)
	}
	logex.Infof("Retrieved product with ID %d: %+v", productId, foundProduct)
}

func updateProduct(product *model.Product) {
	product.Code = newCode
	product.Price = newPrice
	var err error
	product, err = db.UpdateProduct(product)
	if err != nil {
		logex.Fatal(err)
	}
	logex.Infof("Product successfully updated: %+v", product)
}

func deleteProduct(productId uint) {
	err := db.DeleteProduct(productId)
	if err != nil {
		logex.Fatal(err)
	}
	logex.Infof("Product with ID %d successfully deleted", productId)
}

func getAllProducts() {
	products, err := db.GetProducts()
	if err != nil {
		logex.Fatal(err)
	}
	logex.Infof("Found %d products", len(products))
}
