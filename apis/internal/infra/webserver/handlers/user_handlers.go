package handlers

import (
	"encoding/json"
	"net/http"
	"time"

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

func NewUserHandler(userDB database.UserInterface, Jwt *jwtauth.JWTAuth, JwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        userDB,
		Jwt:           Jwt,
		JwtExperiesIn: JwtExperiesIn,
	}
}

func (userHandler *UserHandler) GetJWT(response http.ResponseWriter, request *http.Request) {
	var userDTO dto.GetJWTInput

	err := json.NewDecoder(request.Body).Decode(&userDTO)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := userHandler.UserDB.FindByEmail(userDTO.Email)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !user.ValidatePassword(userDTO.Password) {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := userHandler.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(userHandler.JwtExperiesIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token`
	}{
		AccessToken: token,
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(accessToken)
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
