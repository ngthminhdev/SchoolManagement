package router

import (
	"net/http"

	"GolangBackend/internal/handler"
	"GolangBackend/internal/middleware"
)

func BuildRoutes() http.Handler {
	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/", middleware.HttpLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.HeathCheck(w, r)
		case http.MethodPost:
			handler.Hello(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))
	return mux
}
