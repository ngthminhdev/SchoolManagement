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
	Code    string `json:"code"`
	Data    any    `json:"data"`
}

type WhileListPath struct {
	Method string
	Path   string
}
