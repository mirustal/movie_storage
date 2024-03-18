package models

import "net/http"

type Handlers interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
	SearchMovies(w http.ResponseWriter, r *http.Request)
	AddMovie(w http.ResponseWriter, r *http.Request)
	DeleteMovie(w http.ResponseWriter, r *http.Request)
	UpdateMovie(w http.ResponseWriter, r *http.Request)
	GetActorMovies(w http.ResponseWriter, r *http.Request)
	AddActor(w http.ResponseWriter, r *http.Request)
	DeleteActor(w http.ResponseWriter, r *http.Request)
	UpdateActor(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	RegisterUser(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}