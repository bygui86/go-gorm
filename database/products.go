package database

import (
	"errors"
	"github.com/bygui86/go-gorm/model"
	"gopkg.in/logex.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *DbInterfaceImpl) GetProducts() ([]*model.Product, error) {
	logex.Info("Get all products")
	var products []*model.Product
	err := d.db.Preload(clause.Associations).Find(&products).Error
	return products, err
}

func (d *DbInterfaceImpl) GetProductById(productId uint) (*model.Product, error) {
	logex.Infof("Get product with ID: %d", productId)
	var product model.Product

	// without eager loading
	//err := d.db.First(&product, productId).Error

	// with eager loading of all associations
	err := d.db.Preload(clause.Associations).First(&product, productId).Error

	// with eager loading of a single association
	//err := d.db.Preload("Producers").First(&product, productId).Error

	// with eager loading using inner join for a single association
	//err := d.db.Joins("Producers").First(&product, productId).Error

	return &product, err
}

func (d *DbInterfaceImpl) CreateProduct(product *model.Product) (*model.Product, error) {
	logex.Infof("Create product: %+v", product)
	err := d.db.Create(product).Error
	return product, err
}

/*
	UpdateProduct uses non-zero fields struct to update the product.
	The input product must have the ID, otherwise it's not possible to update it.
*/
func (d *DbInterfaceImpl) UpdateProduct(updatedProduct *model.Product) (*model.Product, error) {
	logex.Infof("Update product: %+v", updatedProduct)
	if updatedProduct.ID == 0 {
		return updatedProduct, errors.New("product ID not valid")
	}
	product := &model.Product{
		Model: gorm.Model{
			ID: updatedProduct.ID,
		},
	}
	err := d.db.Model(&product).Updates(updatedProduct).Error
	return product, err
}

func (d *DbInterfaceImpl) SoftDeleteProduct(productId uint) error {
	logex.Infof("Soft delete product with ID: %d", productId)
	return d.db.Delete(&model.Product{}, productId).Error
}

func (d *DbInterfaceImpl) DeleteProduct(productId uint) error {
	logex.Infof("Delete product with ID: %d", productId)
	return d.db.Unscoped().Delete(&model.Product{}, productId).Error
}
