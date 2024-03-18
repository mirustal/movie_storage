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


	mux.Handle("POST /actors", http.HandlerFunc(handler.AddActor))
	mux.Handle("POST /register", http.HandlerFunc(handler.RegisterUser))
	mux.Handle("POST /login", http.HandlerFunc(handler.LoginUser))
	mux.Handle("GET /actors/{actorId}/movies", http.HandlerFunc(handler.GetActorMovies))
	mux.Handle("DELETE /actors/{actorId}", http.HandlerFunc(handler.DeleteActor))
	mux.Handle("PATCH /actors/{actorId}", http.HandlerFunc(handler.UpdateActor))
	mux.Handle("POST /movies", http.HandlerFunc(handler.AddMovie))
	mux.Handle("GET /movies", http.HandlerFunc(handler.GetMovies))
	mux.Handle("PATCH /movies/{movieId}", http.HandlerFunc(handler.UpdateMovie))
	mux.Handle("DELETE /movies/{movieId}", http.HandlerFunc(handler.DeleteMovie))



	return &Router{
		mux: mux,
	}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
