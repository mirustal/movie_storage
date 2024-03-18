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
	var body models.ActorsIdRequest
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(body)
	answer, err := han.handler.GetActorsMovies(context.TODO(), body)
	if err != nil {
		fmt.Printf("error get actor: %v\n", err)
		http.Error(w, fmt.Sprintf("Error get actor: %v", err), http.StatusBadRequest)
		return
	}
	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)


}


func (han *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	actorID := r.PathValue("actorId")
	var body models.ActorResponse
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}

	answer, err := han.handler.UpdateActor(context.TODO(), actorID, body)
	if err != nil {
		fmt.Printf("error updating actor: %v\n", err)
		http.Error(w, fmt.Sprintf("Error updating actor: %v", err), http.StatusBadRequest)
		return
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}
// ActorsPost - Добавление актёра
func (han *Handler) AddActor(w http.ResponseWriter, r *http.Request) {
	var body models.ActorResponse
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}
	
	answer, err := han.handler.AddActor(context.TODO(), models.ActorResponse{
		Name: body.Name,
		Gender: body.Gender,
		BirthDate: body.BirthDate,
	})
	if err != nil {
		fmt.Printf("error add actor: %v\n", err)
		http.Error(w, fmt.Sprintf("Error add actor: %v", err), http.StatusBadRequest)
		return
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)

}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (han *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	typeGet := r.PathValue("typeGet")
	answer, err := han.handler.GetMovies(context.TODO(), typeGet)
	if err != nil {
		fmt.Printf("error get movie: %v\n", err)
		http.Error(w, fmt.Sprintf("Error get movie: %v", err), http.StatusBadRequest)
		return
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (han *Handler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	typeGet := r.PathValue("search")
	answer, err := han.handler.SearchMovies(context.TODO(), typeGet)
	if err != nil {
		fmt.Printf("error search movie: %v\n", err)
		http.Error(w, fmt.Sprintf("Error search movie: %v", err), http.StatusBadRequest)
		return
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
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
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}
	answer, err := han.handler.UpdateMovie(context.TODO(), movieId, models.MovieResponse{
		Title: body.Title,
		Description: body.Description,
		ReleaseDate: body.ReleaseDate,
		Rating: body.Rating,
	})
	if err != nil {
		fmt.Printf("error update actor: %v\n", err)
		http.Error(w, fmt.Sprintf("Error update actor: %v", err), http.StatusBadRequest)
		return
	}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

// MoviesPost - Добавление фильма
func (han *Handler) AddMovie(w http.ResponseWriter, r *http.Request) {
	var body models.MovieResponse
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}

	if len(body.Actors) == 0 {
		fmt.Printf("Actors can't be empty\n")
		http.Error(w, fmt.Sprintf("Actors can't be empty"), http.StatusBadRequest)
		return
	}
	answer, err := han.handler.AddMovie(context.TODO(), models.MovieResponse{
		Title: body.Title,
		Description: body.Description,
		ReleaseDate: body.ReleaseDate,
		Rating: body.Rating,
		Actors: body.Actors,
	})
	if err != nil {
		fmt.Printf("error add movie: %v\n", err)
		http.Error(w, fmt.Sprintf("Error updating movie: %v", err), http.StatusBadRequest)
		return
		}

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}

func (han *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
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
		if err != nil {
			fmt.Printf("error login user: %v\n", err)
			http.Error(w, fmt.Sprintf("Error login user: %v", err), http.StatusBadRequest)
			return
			}
	}

	acs, err := utils.CreateAccessToken(body.Username, body.Password)
	if err != nil {
		fmt.Printf("error create acecess token: %v\n", err)
		http.Error(w, fmt.Sprintf("Error acces token: %v", err), http.StatusBadRequest)
		return
		}

	utils.SetCookie(w, "accesst", acs)

	EncodeJSONResponse(answer, &GetErrorCode(OkRequest).Code, w)
}


// RegisterPost - Регистрация пользователя и выдача токенов
func (han *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterRequest
	err := DecodeJSONRequest(w, r, &body);
	if err != nil {
		fmt.Printf("error bad json body: %v\n", err)
		http.Error(w, fmt.Sprintf("error bad json body: %v", err), http.StatusBadRequest)
		return
	}
	answer, err := han.handler.RegisterUser(context.TODO(), models.RegisterRequest{
		Username: body.Username,
		Password: body.Password,
		Role: body.Role,
	})
	if err != nil {
		fmt.Printf("error Register user: %v\n", err)
		http.Error(w, fmt.Sprintf("Error register user: %v", err), http.StatusBadRequest)
		return
		}
	acs, err := utils.CreateAccessToken(body.Username, body.Password)
	if err != nil {
		fmt.Printf("error create acecess token: %v\n", err)
		http.Error(w, fmt.Sprintf("Error acces token: %v", err), http.StatusBadRequest)
		return
		}
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

