package controllers

import (
	"GolangBackend/helper"
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BaseController[T entities.IBaseEntity, S services.IBaseService[T]] struct {
	service S
	path    string
}

func NewBaseController[T entities.IBaseEntity, S services.IBaseService[T]](
	service S,
	path string,
) *BaseController[T, S] {
	return &BaseController[T, S]{
		service: service,
		path:    path,
	}
}

func (c *BaseController[T, S]) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/"+c.path, c.FindAll).Methods("GET")
	router.HandleFunc("/"+c.path+"/{id}", c.FindById).Methods("GET")
	router.HandleFunc("/"+c.path, c.Create).Methods("POST")

	router.HandleFunc("/"+c.path+"/{id}", c.Update).Methods("PUT")
	router.HandleFunc("/"+c.path+"/{id}", c.Delete).Methods("DELETE")
}

func (c *BaseController[T, S]) JsonResponse(w http.ResponseWriter, data dto.APIResponse) {
	response, err := json.Marshal(data.ToAPIResponse())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Json encoding to response error"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Status)
	w.Write(response)
}

func (c *BaseController[T, S]) ErrorResponse(w http.ResponseWriter, status int, message *string) {
	errorMsg := "An error has occurred"

	if message != nil {
		errorMsg = *message
	}

	data := dto.APIResponse{
		Status: status,
		Error:  errorMsg,
	}

	c.JsonResponse(w, data)
}

func (c *BaseController[T, S]) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	offset, err := strconv.Atoi(query.Get("skip"))
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	fields := []string{}
	selectFields := query.Get("fields")

	if selectFields != "" {
		fields = append(fields, selectFields)
	}

	entities, err := c.service.FindAll(ctx, &dto.ListOptions{
		Offset: offset,
		Limit:  limit,
		Fields: fields,
	})

	if err != nil {
		c.ErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	coverdEntities := make([]map[string]any, len(entities))
	for i, entity := range entities {
		coverdEntity := entity.ToMap()
		delete(coverdEntity, "password")
		coverdEntities[i] = coverdEntity
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    coverdEntities,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}

func (c *BaseController[T, S]) FindById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	query := r.URL.Query()

	fields := []string{}
	selectFields := query.Get("fields")

	if selectFields != "" {
		fields = append(fields, selectFields)
	}

	entity, err := c.service.FindById(ctx, &dto.GetByIdOptions{
		Fields: fields,
		ID:     id,
	})

	if err != nil {
		helper.LogError(err, "FindById Error")
		c.ErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	coverdEntity := entity.ToMap()
	delete(coverdEntity, "password")

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    coverdEntity,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}

func (c *BaseController[T, S]) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var entity T

	fail := func(status int, message string) {
		c.ErrorResponse(w, status, &message)
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&entity)
	if err != nil {
		fail(http.StatusBadRequest, "Create entity error: invalid JSON")
		return
	}

	newEnity, err := c.service.Create(ctx, entity)

	if err != nil {
		fail(http.StatusBadRequest, "Create entity error: service falure")
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    newEnity.ToMap(),
		Message: "OK",
	}

	c.JsonResponse(w, data)
}

func (c *BaseController[T, S]) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	var entity T

	fail := func(status int, message string) {
		c.ErrorResponse(w, status, &message)
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&entity)
	if err != nil {
		fail(http.StatusBadRequest, "Update entity error: invalid JSON")
		return
	}

	newEnitty, err := c.service.Update(ctx, id, entity)

	if err != nil {
		fail(http.StatusBadRequest, "Update entity error: service falure")
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    newEnitty.ToMap(),
		Message: "OK",
	}


	c.JsonResponse(w, data)
}


func (c *BaseController[T, S]) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	isDeleted, err := c.service.Delete(ctx, id)
	if err != nil {
		msg := "Update entity error"
		c.ErrorResponse(w, http.StatusBadRequest, &msg)
		return
	}

	data := dto.APIResponse{
		Status:  http.StatusOK,
		Data:    isDeleted,
		Message: "OK",
	}

	c.JsonResponse(w, data)
}
