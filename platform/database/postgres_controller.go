package database

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"movie_storage/internal/models"
	storage "movie_storage/internal/models/queries"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	query := `
	DELETE FROM actors
	WHERE id = $1;
`

	_, err := db.db.ExecContext(ctx, query, actorId)
	if err != nil {
		return fmt.Errorf("error executing delete actor query: %w", err)
	}

	return nil
}

// ActorsActorIdMoviesGet - Получение списка фильмов с участнием актера
func (db *API) GetActorMovies(ctx context.Context, actorId string) ([]models.Movie, error) {

	return nil, nil
}

// ActorsActorIdPatch - Изменение информации об актёре
func (db *API) UpdateActor(ctx context.Context, actorId string, actor models.ActorResponse) (models.Actor, error) {
    var queryParams []interface{}
    updateParts := []string{}

    if actor.Name != "" {
        queryParams = append(queryParams, actor.Name)
        updateParts = append(updateParts, fmt.Sprintf("name = $%d", len(queryParams)))
    }
    if actor.Gender != "" {
        queryParams = append(queryParams, actor.Gender)
        updateParts = append(updateParts, fmt.Sprintf("gender = $%d", len(queryParams)))
    }
    if actor.BirthDate != "" {
        queryParams = append(queryParams, actor.BirthDate)
        updateParts = append(updateParts, fmt.Sprintf("birth_date = $%d", len(queryParams)))
    }

    if len(updateParts) == 0 {
        return models.Actor{}, fmt.Errorf("no update data provided")
    }

    queryParams = append(queryParams, actorId)
    query := fmt.Sprintf(`
        UPDATE actors
        SET %s
        WHERE id = $%d
        RETURNING id, name, gender, birth_date;
    `, strings.Join(updateParts, ", "), len(queryParams))

    var updatedActor models.Actor
	err := db.db.QueryRowContext(ctx, query, queryParams...).Scan(&updatedActor.Id, &updatedActor.Name, &updatedActor.Gender, &updatedActor.BirthDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Actor{}, fmt.Errorf("update failed: actor with ID %s not found", actorId)
		}

		return models.Actor{}, fmt.Errorf("update failed: %w", err)
	}
	
	return updatedActor, nil
}

// ActorsPost - Добавление актёра
func (db *API) AddActor(ctx context.Context, actor models.ActorResponse) (models.Actor, error) {
    query := `
        INSERT INTO actors (name, gender, birth_date)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

    var actorID string
    err := db.db.QueryRowContext(ctx, query, actor.Name, actor.Gender, actor.BirthDate).Scan(&actorID)
    if err != nil {
        return models.Actor{}, fmt.Errorf("failed to add actor: %w", err)
    }

    return models.Actor{
        Id: actorID,
        Name: actor.Name,
        Gender: actor.Gender,
        BirthDate: actor.BirthDate,
    }, nil
}


// MoviesGet - Получение списка фильмов с сортировкой и поиском
func (db *API) GetMovies(ctx context.Context, sort string, order string, title string, actorName string) ([]models.Movie, error) {

	return []models.Movie{}, nil
}

// MoviesMovieIdDelete - Удаление фильма
func (db *API) DeleteMovie(ctx context.Context, movieId string) (error) {
	query := `
	DELETE FROM movies
	WHERE id = $1;
`
	err := db.db.QueryRowContext(ctx, query, movieId)
	if err != nil {

	}
	return nil
}

// MoviesMovieIdPatch - Частичное обновление информации о фильме
func (db *API) UpdateMovie(ctx context.Context, movieId string, movie models.MovieResponse) (models.Movie, error) {
    var queryParams []interface{}
    updateParts := []string{}
	movieRating, err := strconv.ParseFloat(movie.Rating, 2)
	if err != nil {
		return models.Movie{}, err
	}

    if movie.Title != "" {
        queryParams = append(queryParams, movie.Title)
        updateParts = append(updateParts, fmt.Sprintf("title = $%d", len(queryParams)))
    }
    if movie.Description != "" {
        queryParams = append(queryParams, movie.Description)
        updateParts = append(updateParts, fmt.Sprintf("description = $%d", len(queryParams))) 
	}
    if movie.ReleaseDate != "" {
        queryParams = append(queryParams, movie.ReleaseDate)
        updateParts = append(updateParts, fmt.Sprintf("release_date = $%d", len(queryParams)))
    }
    if movieRating >= 0 && movieRating <= 10 { 
        queryParams = append(queryParams, movieRating)
        updateParts = append(updateParts, fmt.Sprintf("rating = $%d", len(queryParams)))
    }

    if len(updateParts) == 0 {
        return models.Movie{}, fmt.Errorf("no update data provided")
    }

    queryParams = append(queryParams, movieId)
    query := fmt.Sprintf(`
        UPDATE movies
        SET %s
        WHERE id = $%d
        RETURNING id, title, description, release_date, rating;
    `, strings.Join(updateParts, ", "), len(queryParams))

    var updatedMovie models.Movie
    err = db.db.QueryRowContext(ctx, query, queryParams...).Scan(&updatedMovie.Id, &updatedMovie.Title, &updatedMovie.Description, &updatedMovie.ReleaseDate, &updatedMovie.Rating)
    if err != nil {
        return models.Movie{}, err
    }

    return updatedMovie, nil
}

// MoviesPost - Добавление фильма
func (db *API) AddMovie(ctx context.Context, movie models.MovieResponse) (models.Movie, error) {
	query := `
	INSERT INTO movies (title, description, release_date, rating)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
`

	var userID string
	err := db.db.QueryRowContext(ctx, query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).Scan(&userID)
	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{
		Id: userID,
		Title: movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Rating: movie.Rating,
	}, nil
}

func (db *API) LoginUser(ctx context.Context, userID string, user models.LoginRequest) (models.UserResponse, error) {
	query := `
	UPDATE users
	SET refresh_token = $1
	WHERE id = $2
	RETURNING id;
	`

	randomToken := time.Now().UTC().GoString() + " "
	refreshToken, err := bcrypt.GenerateFromPassword([]byte(randomToken), bcrypt.DefaultCost)
	if err != nil {
		return models.UserResponse{}, err
	}


	err = db.db.QueryRowContext(ctx, query, string(refreshToken), userID).Scan(&userID)
	if err != nil {
		return models.UserResponse{}, err
	}

	refreshTokenEncoded := base64.StdEncoding.EncodeToString([]byte(randomToken))
	return models.UserResponse{
		UserId: userID,
		RefreshToken: refreshTokenEncoded,
	}, nil
}

// RegisterPost - Регистрация пользователя и выдача токенов
func (db *API) RegisterUser(ctx context.Context, user models.RegisterRequest) (models.UserResponse, error) {
	query := `
		INSERT INTO users (refresh_token, role)
		VALUES ($1, $2)
		RETURNING id;
	`

	randomToken := time.Now().UTC().GoString() + " "
	refreshToken, err := bcrypt.GenerateFromPassword([]byte(randomToken), bcrypt.DefaultCost)
	if err != nil {
		return models.UserResponse{}, err
	}
	if user.Role == "" {
		user.Role = "reader" 
	}
	var userID string
	err = db.db.QueryRowContext(ctx, query, string(refreshToken), user.Role).Scan(&userID)
	if err != nil {
		return models.UserResponse{}, err
	}

	refreshTokenEncoded := base64.StdEncoding.EncodeToString([]byte(randomToken))
	return models.UserResponse{
		UserId: userID,
		RefreshToken: refreshTokenEncoded,
	}, nil
}

// TokenGet - Получение информации о текущем токене
func (db *API) GetToken(ctx context.Context) (string, error) {

	return "", nil
}

// TokenRefreshPost - Обновление токена доступа
func (db *API) RefreshToken(ctx context.Context, refresh string) (string, error) {

	return "", nil
}


