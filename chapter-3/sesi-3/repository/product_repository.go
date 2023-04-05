package repository

import "chapter-3/sesi-3/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
