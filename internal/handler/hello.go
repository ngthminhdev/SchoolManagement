package handler

import (
	"GolangBackend/constants"
	"GolangBackend/helper"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/repositories"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
	userRepo := repositories.NewUserRepository()

  newUser, err := userRepo.Create(ctx, &entities.UserEntity{
    Name: "name1",
    Email: "name1@email.com",
    Phone: "0987654321",
    Password: "name1_password",
    Gender: int8(constants.MALE),
  })

	if err != nil {
		helper.LogError(err, "Users create error")
	}

	helper.LogInfo("newUser %v", newUser)
	helper.LogInfo("newUser %v", newUser.ToMap())

	fmt.Fprintln(w, "Hello world")
}

func HeathCheck(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
	body, _ := io.ReadAll(r.Body)

	userRepo := repositories.NewUserRepository()

	data, err := userRepo.FindAll(ctx, nil)
	if err != nil {
		helper.LogError(err, "Users FindAll error")
	}

	helper.LogInfo("%v", data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)

	fmt.Fprintf(w, "%s", string(body))
}
