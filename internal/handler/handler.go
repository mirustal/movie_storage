package handler

import (
	
	"fmt"
	"movie_storage/internal/models/queries"
	"net/http"
)


type Handler struct {
	handler queries.Storage
}

func NewHandler(api queries.Storage) *Handler {
	return &Handler{
		handler: api,
	}
}

// ActorsActorIdDelete - Удаление актёра
func (han *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {


}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (han *Handler) GetActorMovies(w http.ResponseWriter, r *http.Request) {


}

// ActorsActorIdPatch - Изменение информации об актёре
func (han *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {


}

// ActorsPost - Добавление актёра
func (han *Handler) AddActor(w http.ResponseWriter, r *http.Request) {


}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (han *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {

}

// MoviesMovieIdDelete - Удаление фильма
func (han *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {

}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (han *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

// MoviesPost - Добавление фильма
func (han *Handler) AddMovie(w http.ResponseWriter, r *http.Request) {

}

// RegisterPost - Регистрация пользователя и выдача токенов
func (han *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("типа зарегистрировал")

}

// TokenGet - Получение информации о текущем токене
func (han *Handler) GetToken(w http.ResponseWriter, r *http.Request) {


}

// TokenRefreshPost - Обновление токена доступа
func (han *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

