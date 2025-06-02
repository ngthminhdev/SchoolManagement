package middleware

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/global"
	"GolangBackend/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if global.WhileListPaths[r.RequestURI] != "" {
			if global.WhileListPaths[r.RequestURI] == r.Method {
				next.ServeHTTP(w, r)
				return
			}
		}

		errorResponse := func(msg string) {
			apiResponse := dto.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: msg,
			}

			response, err := json.Marshal(apiResponse)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Json encoding to response error"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
		}

		bearToken := r.Header.Get("Authorization")

		if bearToken == "" {
			errorResponse("Un authorization")
			return
		}

		jwtToken := strings.Split(bearToken, " ")[1]

		isVaidJWT, err := services.VerifyJWT(jwtToken)
		if err != nil {
			errMsg := err.Error()
			errorResponse(errMsg)
			return
		}

		if !isVaidJWT {
			errorResponse("jwt invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
