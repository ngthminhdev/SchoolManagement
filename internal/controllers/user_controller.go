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
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body dto.RegisterDTO

	fail := func(status int, message string) {
		c.ErrorResponse(w, status, &message)
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		fail(http.StatusBadRequest, "Invalid JSON")
		return
	}

	newUser, err := c.service.Register(ctx, &body)

	if err != nil {
		fail(http.StatusBadRequest, err.Error())
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    newUser,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}
