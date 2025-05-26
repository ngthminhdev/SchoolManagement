package dto

type CreateUserDTO struct{}

type RegisterDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   int8   `json:"gender"`
}
