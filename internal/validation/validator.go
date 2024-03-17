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

	return nil
}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (db *ValidationService) GetActorMovies(ctx context.Context, actorId string) ([]models.Movie, error) {

	return nil, nil
}

// ActorsActorIdPatch - Изменение информации об актёре
func (db *ValidationService) UpdateActor(ctx context.Context, actorId string, actor models.Actor) (models.Actor, error) {

	return models.Actor{}, nil
}

// ActorsPost - Добавление актёра
func (db *ValidationService) AddActor(ctx context.Context, actor models.Actor) (models.Actor, error) {

	return models.Actor{}, nil
}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (db *ValidationService) GetMovies(ctx context.Context, sort string, order string, title string, actorName string) ([]models.Movie, error) {

	return []models.Movie{}, nil
}

// MoviesMovieIdDelete - Удаление фильма
func (db *ValidationService) DeleteMovie(ctx context.Context, movieId string) (error) {

	return nil
}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (db *ValidationService) UpdateMovie(ctx context.Context, movieId string, movie models.Movie) (models.Movie, error) {

	return models.Movie{}, nil
}

// MoviesPost - Добавление фильма
func (db *ValidationService) AddMovie(ctx context.Context, movie models.Movie) (models.Movie, error) {

	return models.Movie{}, nil
}

// RegisterPost - Регистрация пользователя и выдача токенов
func (db *ValidationService) RegisterUser(ctx context.Context, user models.RegisterRequest) (string, error) {
	fmt.Println("типа зарегистрировал")
	fmt.Println(user)
	return "", nil
}

// TokenGet - Получение информации о текущем токене
func (db *ValidationService) GetToken(ctx context.Context) (string, error) {

	return "", nil
}

// TokenRefreshPost - Обновление токена доступа
func (db *ValidationService) RefreshToken(ctx context.Context, refresh string) (string, error) {

	return "", nil
}



