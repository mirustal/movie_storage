package validation

import (
	"context"
	"errors"
	"fmt"
	"movie_storage/internal/models"
	"movie_storage/internal/models/queries"
)


var (
	ErrIdInvalid = errors.New("invalid id")
)

type ValidationService struct {
	requestApi queries.Storage
}

func NewValidator(api queries.Storage) *ValidationService {
	return &ValidationService{
		requestApi: api,
	}
}


// ActorsActorIdDelete - Удаление актёра
func (db *ValidationService) DeleteActor(ctx context.Context, actorId string) (error) {
	err := db.requestApi.DeleteActor(ctx, actorId)
	if err != nil {
		return err
	}
	return nil
}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (db *ValidationService) GetActorsMovies(ctx context.Context, actorId models.ActorsIdRequest) ([]models.ActorWithMovies, error) {
	ans, err := db.requestApi.GetActorsMovies(ctx, actorId)
	if err != nil {
		return nil, err
	}
	return ans, nil
}

// ActorsActorIdPatch - Изменение информации об актёре
func (db *ValidationService) UpdateActor(ctx context.Context, actorId string, actor models.ActorResponse) (models.Actor, error) {
	ans, err := db.requestApi.UpdateActor(ctx, actorId, actor)
	if err != nil {
		return models.Actor{}, err
	}
	return ans, nil
}

// ActorsPost - Добавление актёра
func (db *ValidationService) AddActor(ctx context.Context, actor models.ActorResponse) (models.Actor, error) {
	ans, err := db.requestApi.AddActor(ctx, actor)
	if err != nil {
		return models.Actor{}, err
	}
	return ans, nil
}



// MoviesGet - Получение списка фильмов 
func (db *ValidationService) GetMovies(ctx context.Context, typeGet string) ([]models.MovieResponseActor, error) {
	if typeGet != "title" && typeGet != "date" && typeGet != "rating" && typeGet != "" {
        return nil, fmt.Errorf("This type of receipt is not suitable ")
    }
	ans, err := db.requestApi.GetMovies(ctx, typeGet)
	if err != nil {
		return nil, err
	}
	return ans, nil
}

// MoviesSearch - Получение списка фильмов с сортировкой и поиском
func (db *ValidationService) SearchMovies(ctx context.Context, search string) ([]models.MovieResponseActor, error) {
	ans, err := db.requestApi.SearchMovies(ctx, search)
	if err != nil {
		return nil, err
	}
	return ans, nil
}

// MoviesMovieIdDelete - Удаление фильма
func (db *ValidationService) DeleteMovie(ctx context.Context, movieId string) (error) {
	err := db.requestApi.DeleteMovie(ctx, movieId)
	if err != nil {
		return err
	}
	return nil
}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (db *ValidationService) UpdateMovie(ctx context.Context, movieId string, movie models.MovieResponse) (models.Movie, error) {
	ans, err := db.requestApi.UpdateMovie(ctx, movieId, movie)
	if err != nil {
		return models.Movie{}, err
	}
	return ans, nil
}

// MoviesPost - Добавление фильма
func (db *ValidationService) AddMovie(ctx context.Context, movie models.MovieResponse) (models.MovieResponseActor, error) {
	ans, err := db.requestApi.AddMovie(ctx, movie)
	if err != nil {
		return models.MovieResponseActor{}, err
	}
	return ans, nil
}

func (db *ValidationService) LoginUser(ctx context.Context, userId string, user models.LoginRequest) (models.UserResponse, error) {
	ans, err := db.requestApi.LoginUser(ctx, userId, user)
	if err != nil {
		return models.UserResponse{}, err
	}
	return ans, nil
}

// RegisterPost - Регистрация пользователя и выдача токенов
func (db *ValidationService) RegisterUser(ctx context.Context, user models.RegisterRequest) (models.UserResponse, error) {
	if user.Username == "" {
        return models.UserResponse{}, fmt.Errorf("username is required")
    }
    if user.Password == "" {
        return models.UserResponse{}, fmt.Errorf("password is required")
    }

	ans, err := db.requestApi.RegisterUser(ctx, user)
	if err != nil {
		return models.UserResponse{}, err
	}
	return ans, nil
}

// TokenGet - Получение информации о текущем токене
func (db *ValidationService) GetToken(ctx context.Context) (string, error) {

	return "", nil
}

// TokenRefreshPost - Обновление токена доступа
func (db *ValidationService) RefreshToken(ctx context.Context, refresh string) (string, error) {

	return "", nil
}



