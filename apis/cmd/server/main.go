package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lccoronel/golang-full-cycle/apis/configs"
	_ "github.com/lccoronel/golang-full-cycle/apis/docs"
	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
	"github.com/lccoronel/golang-full-cycle/apis/internal/infra/database"
	"github.com/lccoronel/golang-full-cycle/apis/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title			Go Expert API Example
// @version			1.0
// @description		Product API with authentication
// @termsOfService	http://swagger.io/terms/

// @contact.name	Lucas Coronel
// @contact.url		https://github.com/lccoronel
// @contact.email	lucascoronel0597@gmail.com

// @license.name	Full Cycle License

// @host			localhost:8000
// @Basepath		/
// @securityDefinitions.apikey ApiKeyAuth
// @in 				header
// @name 			Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userhandler := handlers.NewUserHandler(userDB)

	router := chi.NewRouter()
	// router.Use(LogRequest)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("JWTExperiesIn", config.JWTExpiresIn))

	router.Route("/products", func(productsRouter chi.Router) {
		productsRouter.Use(jwtauth.Verifier(config.TokenAuth))
		productsRouter.Use(jwtauth.Authenticator)

		productsRouter.Post("/", productHandler.CreateProduct)
		productsRouter.Get("/{id}", productHandler.GetProduct)
		productsRouter.Get("/", productHandler.GetAllProducts)
		productsRouter.Put("/{id}", productHandler.UpdateProduct)
		productsRouter.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userhandler.CreateUser)
	router.Post("/users/generate_token", userhandler.GetJWT)
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", router)
}

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
// 		log.Printf("Request: %s %s", request.Method, request.URL.Path)
// 		next.ServeHTTP(response, request)
// 	})
// }
