package repository

import (
	"chapter-3/challange-3/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)

	return &product
}

func (repo *ProductRepositoryMock) FindAll() []*models.Product {
	args := repo.Mock.Called()
	return args.Get(0).([]*models.Product)
}
