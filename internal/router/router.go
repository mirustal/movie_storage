package router

import (
	"movie_storage/internal/middleware"
	"movie_storage/internal/models"
	"movie_storage/pkg/configs"

	"net/http"
)



type Router struct {
	mux *http.ServeMux
}

func NewRouter(cfg *configs.Config, handler models.Handlers) *Router {

	mux := http.NewServeMux()

	mux.Handle("POST /actors", middleware.AuthOnlyToken(http.HandlerFunc(handler.AddActor)))
	mux.Handle("POST /register", http.HandlerFunc(handler.RegisterUser))
	mux.Handle("POST /login", http.HandlerFunc(handler.LoginUser))
	mux.Handle("GET /actors/", http.HandlerFunc(handler.GetActorMovies))
	mux.Handle("DELETE /actors/{actorId}", middleware.AuthOnlyToken(http.HandlerFunc(handler.DeleteActor)))
	mux.Handle("PATCH /actors/{actorId}", middleware.AuthOnlyToken(http.HandlerFunc(handler.UpdateActor)))
	mux.Handle("POST /movies", middleware.AuthOnlyToken(http.HandlerFunc(handler.AddMovie)))
	mux.Handle("GET /movies/{typeGet}", middleware.AuthOnlyToken(http.HandlerFunc(handler.GetMovies)))
	mux.Handle("GET /movies/", middleware.AuthOnlyToken(http.HandlerFunc(handler.GetMovies)))
	mux.Handle("GET /movies/search/{search}", middleware.AuthOnlyToken(http.HandlerFunc(handler.SearchMovies)))
	mux.Handle("PATCH /movies/{movieId}", middleware.AuthOnlyToken(http.HandlerFunc(handler.UpdateMovie)))
	mux.Handle("DELETE /movies/{movieId}", middleware.AuthOnlyToken(http.HandlerFunc(handler.DeleteMovie)))

	return &Router{
		mux: mux,
	}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
