package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lccoronel/golang-full-cycle/apis/internal/dto"
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
	"github.com/lccoronel/golang-full-cycle/apis/internal/infra/database"
	entityPKG "github.com/lccoronel/golang-full-cycle/apis/pkg/entity"
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

func (productHandler *ProductHandler) GetProduct(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := productHandler.ProductDB.FindByID(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}

func (productHandler *ProductHandler) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPKG.ParseID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = productHandler.ProductDB.FindByID(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	err = productHandler.ProductDB.Update(&product)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}
