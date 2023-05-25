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

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (userHandler *UserHandler) GetJWT(response http.ResponseWriter, request *http.Request) {
	jwt := request.Context().Value("jwt").(*jwtauth.JWTAuth)
	JWTExperiesIn := request.Context().Value("JWTExperiesIn").(int)
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

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(JWTExperiesIn)).Unix(),
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

// Create user godoc
// @Summary			Create user
// @Description		Create user
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			request		body		dto.CreateUserInput 	true	"user request"
// @Success 		201
// @Failure 		500			{object}	Error
// @router			/users 		[post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(response).Encode(error)
		return
	}

	err = userHandler.UserDB.Create(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(response).Encode(error)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
