package dto

type SelectFieldsOptions struct {
  Fields []string
}

type ListOptions struct {
  Fields []string
  Offset int
  Limit int
}

type GetByIdOptions struct {
  Fields []string
  ID string
}

type APIResponse struct {
  Status int
  Message string
  Data any
  Error string
}
