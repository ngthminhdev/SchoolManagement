package controllers

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	*BaseController[*entities.UserEntity, services.IUserService]
	service services.IUserService
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{
		BaseController: NewBaseController(
			service,
			"users",
		),
		service: service,
	}
}

func (c *UserController) RegisterRoutes(router *mux.Router) {
	c.BaseController.RegisterRoutes(router)
	router.HandleFunc("/"+c.path+"/register", c.Register).Methods("POST")
	router.HandleFunc("/"+c.path+"/login", c.Login).Methods("POST")
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body dto.RegisterDTO

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		errMsg := err.Error()
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	newUser, err := c.service.Register(ctx, &body)
	if err != nil {
		errMsg := err.Error()
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    newUser,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body dto.LoginDTO

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		errMsg := err.Error()
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	if body.Account == "" {
		errMsg := "account is required"
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	if body.Password == "" {
		errMsg := "password is required"
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	newUser, err := c.service.Login(ctx, &body)

	if err != nil {
		errMsg := err.Error()
		c.ErrorResponse(w, http.StatusBadRequest, &errMsg)
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    newUser,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}
