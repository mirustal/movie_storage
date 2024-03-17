package router

import (

	"movie_storage/internal/models"
	"movie_storage/pkg/configs"

	"net/http"
)



type Router struct {
	mux *http.ServeMux
}

func NewRouter(cfg *configs.Config, handler models.Handlers) *Router {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	mux.Handle("POST /register", http.HandlerFunc(handler.RegisterUser))

	return &Router{
		mux: mux,
	}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
