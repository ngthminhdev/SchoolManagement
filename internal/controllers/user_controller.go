package controllers

import (
	"GolangBackend/internal/entities"
	"GolangBackend/internal/services"

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
}
