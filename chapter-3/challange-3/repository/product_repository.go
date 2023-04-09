package repository

import "chapter-3/challange-3/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
