package service

import (
	"chapter-3/challange-3/models"
	"chapter-3/challange-3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "2").Return(nil)

	product, err := productService.GetOneProduct("2")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.Product{
		Title:       "Buku A",
		Description: "ini sebuah buku",
	}

	productRepository.Mock.On("FindById", "3").Return(product)

	result, err := productService.GetOneProduct("3")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Title, result.Title, "result has to be '3'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'ini sebuah buku'")
	assert.Equal(t, &product, result, "result has to be a product data with id '3'")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(make([]*models.Product, 0))

	products, err := productService.GetAllProduct()

	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "no products found", err.Error(), "error response has to be 'no products found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	products := []*models.Product{
		{
			Title:       "Buku A",
			Description: "ini sebuah buku",
		},
		{
			Title:       "Buku A",
			Description: "ini sebuah buku",
		},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(products), len(result))
}
