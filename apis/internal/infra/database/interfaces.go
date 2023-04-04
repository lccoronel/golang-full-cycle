package database

import (
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
