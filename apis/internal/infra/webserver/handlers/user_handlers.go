package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lccoronel/golang-full-cycle/apis/internal/dto"
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
	"github.com/lccoronel/golang-full-cycle/apis/internal/infra/database"
)

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (userHandler *UserHandler) CreateUser(response http.ResponseWriter, request *http.Request) {
	var userDTO dto.CreateUserInput

	err := json.NewDecoder(request.Body).Decode(&userDTO)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = userHandler.UserDB.Create(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
