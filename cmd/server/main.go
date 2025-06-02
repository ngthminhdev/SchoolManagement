package main

import (
	"log"
	"net/http"

	"GolangBackend/config"
	"GolangBackend/helper"
	"GolangBackend/internal/controllers"
	"GolangBackend/internal/db"
	"GolangBackend/internal/middleware"
	"GolangBackend/internal/repositories"
	"GolangBackend/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	env := config.GetEnv("ENV", "development")
	config.SetWhiteListPaths()

	helper.InitLogger(env == "production")
	helper.LogInfo("Running with ENV [%s]\n", env)

	db.ConnectDatabase()

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userController.RegisterRoutes(apiRouter)

	// Middleware
	router.Use(middleware.HttpLog)
	router.Use(middleware.JWTAuth)

	var port string = config.GetPort()
	var server http.Server = http.Server{
		Addr:    port,
		Handler: router,
	}

	helper.LogInfo("Server is running at http://localhost%s\n", port)

	var err error = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
