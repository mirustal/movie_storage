package database

import (
	"context"
	"errors"
	"fmt"
	"movie_storage/internal/models"
	storage "movie_storage/internal/models/queries"
)

var (
	ErrFilmNotExist   = errors.New("actor does not exist")
	ErrEmptyUpdate    = errors.New("no updates to apply")
	ErrFilmActorExist = errors.New("given film and actor are already bound")
	ErrActorNotExist  = errors.New("actor with given id does not exist")
	ErrZeroActors     = errors.New("no actors affected")
)

var _ storage.Storage = (*API)(nil)

type API struct {
	db Client
}

func NewAPI(db Client) *API {
	return &API{
		db: db,
	}
}


// ActorsActorIdDelete - Удаление актёра
func (db *API) DeleteActor(ctx context.Context, actorId string) (error) {

	return nil
}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (db *API) GetActorMovies(ctx context.Context, actorId string) ([]models.Movie, error) {

	return nil, nil
}

// ActorsActorIdPatch - Изменение информации об актёре
func (db *API) UpdateActor(ctx context.Context, actorId string, actor models.Actor) (models.Actor, error) {

	return models.Actor{}, nil
}

// ActorsPost - Добавление актёра
func (db *API) AddActor(ctx context.Context, actor models.Actor) (models.Actor, error) {

	return models.Actor{}, nil
}

// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (db *API) GetMovies(ctx context.Context, sort string, order string, title string, actorName string) ([]models.Movie, error) {

	return []models.Movie{}, nil
}

// MoviesMovieIdDelete - Удаление фильма
func (db *API) DeleteMovie(ctx context.Context, movieId string) (error) {

	return nil
}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (db *API) UpdateMovie(ctx context.Context, movieId string, movie models.Movie) (models.Movie, error) {

	return models.Movie{}, nil
}

// MoviesPost - Добавление фильма
func (db *API) AddMovie(ctx context.Context, movie models.Movie) (models.Movie, error) {

	return models.Movie{}, nil
}

// RegisterPost - Регистрация пользователя и выдача токенов
func (db *API) RegisterUser(ctx context.Context, user models.RegisterRequest) (string, error) {
	fmt.Println("типа зарегистрировал")
	fmt.Println(user)
	return "", nil
}

// TokenGet - Получение информации о текущем токене
func (db *API) GetToken(ctx context.Context) (string, error) {

	return "", nil
}

// TokenRefreshPost - Обновление токена доступа
func (db *API) RefreshToken(ctx context.Context, refresh string) (string, error) {

	return "", nil
}


