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
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func (a *APIResponse) ToAPIResponse() map[string]any {
	m := map[string]any{}

	m["status"] = a.Status
	m["message"] = a.Message
	m["data"] = a.Data
	m["error"] = a.Error

	return m
}
