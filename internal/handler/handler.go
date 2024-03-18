package handler

import (
	"context"
	"fmt"
	"log"

	"movie_storage/internal/models"
	"movie_storage/internal/models/queries"
	"movie_storage/pkg/utils"
	. "movie_storage/pkg/utils"

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
	actorID := r.PathValue("actorId")
	err := han.handler.DeleteActor(context.TODO(), actorID)
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (han *Handler) GetActorMovies(w http.ResponseWriter, r *http.Request) {

}

// ActorsActorIdPatch - Изменение информации об актёре
func (han *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	actorID := r.PathValue("actorId")
	var body models.ActorResponse
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return 
	}

	answer, err := han.handler.UpdateActor(context.TODO(), actorID, body)
	if err != nil {
    fmt.Printf("error updating actor: %v\n", err)
    http.Error(w, fmt.Sprintf("Error updating actor: %v", err), http.StatusBadRequest)
	return
	}

	// Если ошибки нет, продолжаем с отправкой успешного ответа
	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}
// ActorsPost - Добавление актёра
func (han *Handler) AddActor(w http.ResponseWriter, r *http.Request) {
	var body models.ActorResponse
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return
	}
	
	answer, err := han.handler.AddActor(context.TODO(), models.ActorResponse{
		Name: body.Name,
		Gender: body.Gender,
		BirthDate: body.BirthDate,
	})
	if err != nil {
		log.Println(err)
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)

}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (han *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {

}

// MoviesMovieIdDelete - Удаление фильма
func (han *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("movieId")
	err := han.handler.DeleteMovie(context.TODO(), movieId)
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (han *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("movieId")
	var body models.MovieResponse
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return
	}
	answer, err := han.handler.UpdateMovie(context.TODO(), movieId, models.MovieResponse{
		Title: body.Title,
		Description: body.Description,
		ReleaseDate: body.ReleaseDate,
		Rating: body.Rating,
	})
	if err != nil {
		EncodeJSONResponse(err, &GetErrorCode(BadRequest).Code, w)
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

// MoviesPost - Добавление фильма
func (han *Handler) AddMovie(w http.ResponseWriter, r *http.Request) {
	var body models.MovieResponse
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return
	}
	
	answer, err := han.handler.AddMovie(context.TODO(), models.MovieResponse{
		Title: body.Title,
		Description: body.Description,
		ReleaseDate: body.ReleaseDate,
		Rating: body.Rating,
	})
	if err != nil {
		EncodeJSONResponse(err, &GetErrorCode(BadRequest).Code, w)
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

func (han *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return
	}
	cookieId, err := r.Cookie("userid")
	if err != nil {
		log.Println(err)
	}
	answer, err := han.handler.LoginUser(context.TODO(), cookieId.Value, models.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		log.Println(err)
	}

	acs, err := utils.CreateAccessToken(body.Username, body.Password)
	if err != nil {
		log.Println(err)
	}

	utils.SetCookie(w, "accesst", acs)

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}


// RegisterPost - Регистрация пользователя и выдача токенов
func (han *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterRequest
	if ok := DecodeJSONRequest(w, r, &body); !ok {
		return
	}
	answer, err := han.handler.RegisterUser(context.TODO(), models.RegisterRequest{
		Username: body.Username,
		Password: body.Password,
		Role: body.Role,
	})
	if err != nil {

	}
	acs, err := utils.CreateAccessToken(body.Username, body.Password)
	utils.SetCookie(w, "accesst", acs)
	utils.SetCookie(w, "userid", answer.UserId)
	
	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

// TokenGet - Получение информации о текущем токене
func (han *Handler) GetToken(w http.ResponseWriter, r *http.Request) {


}

// TokenRefreshPost - Обновление токена доступа
func (han *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

