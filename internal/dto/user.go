package dto

type RegisterDTO struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Gender   int16  `json:"gender" db:"gender"`
}

type LoginDTO struct {
	Account  string `json:"account" db:"account"`
	Password string `json:"password" db:"password"`
}
