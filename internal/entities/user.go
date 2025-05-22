package entities

type UserEntity struct {
	BaseEntity
	Name     string
	Email    string
	Phone    string
	Password string
	Gender   int8
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
	if v, ok := data["gender"].(int8); ok {
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

func (u *UserEntity) ToSQLParams() []any {
	return append(u.BaseEntity.ToSQLParams(), u.Name, u.Email, u.Phone, u.Password, u.Gender)
}
