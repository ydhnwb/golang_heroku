package repo

import (
	"github.com/ydhnwb/golang_heroku/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	All(userID string) ([]entity.Product, error)
	InsertProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(productID string) error
	FindOneProductByID(ID string) (entity.Product, error)
	FindAllProduct(userID string) ([]entity.Product, error)
}

type productRepo struct {
	connection *gorm.DB
}

func NewProductRepo(connection *gorm.DB) ProductRepository {
	return &productRepo{
		connection: connection,
	}
}

func (c *productRepo) All(userID string) ([]entity.Product, error) {
	products := []entity.Product{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&products)
	return products, nil
}

func (c *productRepo) InsertProduct(product entity.Product) (entity.Product, error) {
	c.connection.Save(&product)
	c.connection.Preload("User").Find(&product)
	return product, nil
}

func (c *productRepo) UpdateProduct(product entity.Product) (entity.Product, error) {
	c.connection.Save(&product)
	c.connection.Preload("User").Find(&product)
	return product, nil
}

func (c *productRepo) FindOneProductByID(productID string) (entity.Product, error) {
	var product entity.Product
	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
	if res.Error != nil {
		return product, res.Error
	}
	return product, nil
}

func (c *productRepo) FindAllProduct(userID string) ([]entity.Product, error) {
	products := []entity.Product{}
	c.connection.Where("user_id = ?", userID).Find(&products)
	return products, nil
}

func (c *productRepo) DeleteProduct(productID string) error {
	var product entity.Product
	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&product)
	return nil
}
