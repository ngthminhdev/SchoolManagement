package dto

type SelectFieldsOptions struct {
	Fields []string
}

type ListOptions struct {
	Fields []string
	Offset int
	Limit  int
}

type GetByIdOptions struct {
	Fields []string
	ID     string
}

type APIResponse struct {
	Status  int
	Message string
	Data    any
	Error   string
}

func (a *APIResponse) ToAPIResponse() map[string]any {
	m := map[string]any{}

	m["status"] = a.Status
	m["message"] = a.Message
	m["data"] = a.Data
	m["error"] = a.Error

	return m
}
