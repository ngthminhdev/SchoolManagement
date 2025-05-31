package entities

type UserEntity struct {
	BaseEntity
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"-" db:"password"`
	Gender   int16   `json:"gender" db:"gender"`
}

type UserResponse struct {
	UserEntity
	Password     string `json:"-"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *UserEntity) FromMap(data map[string]any) {
	u.BaseEntity.FromMap(data)

	if v, ok := data["name"].(string); ok {
		u.Name = v
	}
	if v, ok := data["email"].(string); ok {
		u.Email = v
	}
	if v, ok := data["phone"].(string); ok {
		u.Phone = v
	}
	if v, ok := data["password"].(string); ok {
		u.Password = v
	}
	if v, ok := data["gender"].(int16); ok {
		u.Gender = v
	}
}

func (u *UserEntity) ToMap() map[string]any {
	m := u.BaseEntity.ToMap()

	m["name"] = u.Name
	m["email"] = u.Email
	m["phone"] = u.Phone
	m["password"] = u.Password
	m["gender"] = u.Gender

	return m
}
