package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lccoronel/golang-full-cycle/apis/internal/dto"
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
	"github.com/lccoronel/golang-full-cycle/apis/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (productHandler *ProductHandler) CreateProduct(response http.ResponseWriter, request *http.Request) {
	var productDTO dto.CreateProductInput

	err := json.NewDecoder(request.Body).Decode(&productDTO)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productDTO.Name, productDTO.Price)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = productHandler.ProductDB.Create(product)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
