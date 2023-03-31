package entity

import (
	"errors"
	"time"

	"github.com/lccoronel/golang-full-cycle/apis/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrorIDIsRequired    = errors.New("ID is required")
	ErrorInvalidID       = errors.New("Invalid ID")
	ErrorInvalidPrice    = errors.New("Invalid Price")
	ErrorNameIsRequired  = errors.New("Name is required")
	ErrorPriceIsRequired = errors.New("Price is required")
)

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (product *Product) Validate() error {
	if product.ID.String() == "" {
		return ErrorIDIsRequired
	}

	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrorInvalidID
	}

	if product.Name == "" {
		return ErrorNameIsRequired
	}

	if product.Price == 0 {
		return ErrorPriceIsRequired
	}

	if product.Price < 0 {
		return ErrorInvalidPrice
	}

	return nil
}
