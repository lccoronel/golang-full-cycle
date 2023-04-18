package database

import (
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(product *entity.Product) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}