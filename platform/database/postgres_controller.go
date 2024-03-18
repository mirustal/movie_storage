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
func (db *API) GetActorsMovies(ctx context.Context, actorIds models.ActorsIdRequest) ([]models.ActorWithMovies, error) {
    actorsWithMovies := []models.ActorWithMovies{}

    for _, actorId := range actorIds.ActorID {
        actorWithMovies := models.ActorWithMovies{}
        
        // Получаем информацию об актёре
        actorQuery := `SELECT id, name FROM actors WHERE id = $1;`
        err := db.db.QueryRowContext(ctx, actorQuery, actorId).Scan(&actorWithMovies.ID, &actorWithMovies.Name)
        if err != nil {
            if err == sql.ErrNoRows {
                continue // Актёр не найден, переходим к следующему ID
            }
            return nil, fmt.Errorf("querying actor info: %v", err)
        }

        // Получаем фильмы актёра
        moviesQuery := `
            SELECT m.title, m.release_date, m.rating 
            FROM movies m
            JOIN actors_movies am ON m.id = am.movie_id
            WHERE am.actor_id = $1;`
        rows, err := db.db.QueryContext(ctx, moviesQuery, actorId)
        if err != nil {
            return nil, fmt.Errorf("querying actor's movies: %v", err)
        }
        defer rows.Close()

        movies := []models.MovieDetail{}
        for rows.Next() {
            var movie models.MovieDetail
            if err := rows.Scan(&movie.Title, &movie.ReleaseDate, &movie.Rating); err != nil {
                return nil, fmt.Errorf("scanning movie: %v", err)
            }
            movies = append(movies, movie)
        }
        
        actorWithMovies.Movies = movies
        actorsWithMovies = append(actorsWithMovies, actorWithMovies)
    }

    return actorsWithMovies, nil
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

	var existingActorID int
	checkQuery := `SELECT id FROM actors WHERE name = $1 AND birth_date = $2 AND gender = $3`
    err := db.db.QueryRowContext(ctx, checkQuery, actor.Name, actor.BirthDate, actor.Gender).Scan(&existingActorID)
    if err != nil {
    } else {
        return models.Actor{}, fmt.Errorf("The actor already exists")
    }


    query := `
        INSERT INTO actors (name, gender, birth_date)
        VALUES ($1, $2, $3)
        RETURNING id;
    `
    var actorID string
    err = db.db.QueryRowContext(ctx, query, actor.Name, actor.Gender, actor.BirthDate).Scan(&actorID)
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
func (db *API) GetMovies(ctx context.Context, typeGet string) ([]models.MovieResponseActor, error) {
    orderBy := "rating DESC" // Сортировка по умолчанию

    // Определяем столбец сортировки в зависимости от входного параметра
    switch typeGet {
    case "title":
        orderBy = "title ASC"
    case "releaseDate":
        orderBy = "release_date ASC"
    case "rating":
        orderBy = "rating DESC"
    }

    query := fmt.Sprintf(`
        SELECT id, title, description, release_date, rating
        FROM movies
        ORDER BY %s;`, orderBy)

    rows, err := db.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movies []models.MovieResponseActor
    for rows.Next() {
        var movie models.MovieResponseActor
        err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
        if err != nil {
            return nil, err
		}
		movie.ReleaseDate = strings.Split(movie.ReleaseDate, ":")[0]
        movies = append(movies, movie)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    for i := range movies {
        var actorIDs []int
        queryActors := `SELECT actor_id FROM actors_movies WHERE movie_id = $1;`
        actorRows, err := db.db.QueryContext(ctx, queryActors, movies[i].Id)
        if err != nil {
            return nil, err
        }

        for actorRows.Next() {
            var actorID int
            if err := actorRows.Scan(&actorID); err != nil {
                actorRows.Close()
                return nil, err
            }
            actorIDs = append(actorIDs, actorID)
        }
        actorRows.Close() // Закрываем соединение после обработки каждого набора строк

        for _, actorID := range actorIDs {
            var actor models.Actor
            actorQuery := `SELECT id, name, birth_date FROM actors WHERE id = $1;`
            if err := db.db.QueryRowContext(ctx, actorQuery, actorID).Scan(&actor.Id, &actor.Name, &actor.BirthDate); err != nil {
                return nil, fmt.Errorf("querying actor info: %v", err)
            }
            movies[i].Actors = append(movies[i].Actors, actor)
        }
    }

    return movies, nil
}


func (db *API) SearchMovies(ctx context.Context, searchQuery string) ([]models.MovieResponseActor, error) {

    query := `
        SELECT DISTINCT m.id, m.title, m.release_date, m.rating
        FROM movies m
        LEFT JOIN actors_movies am ON m.id = am.movie_id
        LEFT JOIN actors a ON am.actor_id = a.id
        WHERE m.title ILIKE $1 OR a.name ILIKE $1
        ORDER BY m.rating DESC;
    `

    searchPattern := "%" + searchQuery + "%"

    rows, err := db.db.QueryContext(ctx, query, searchPattern)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movies []models.MovieResponseActor
    for rows.Next() {
        var movie models.MovieResponseActor
        err := rows.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate, &movie.Rating)
        if err != nil {
            return nil, err
        }

        actorsQuery := `
            SELECT a.id, a.name, a.gender, a.birth_date
            FROM actors a
            JOIN actors_movies am ON a.id = am.actor_id
            WHERE am.movie_id = $1;
        `
        actorRows, actorErr := db.db.QueryContext(ctx, actorsQuery, movie.Id)
        if actorErr != nil {
            return nil, actorErr
        }

        var actors []models.Actor
        for actorRows.Next() {
            var actor models.Actor
            if err := actorRows.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.BirthDate); err != nil {
                actorRows.Close() 
                return nil, err
            }
            actors = append(actors, actor)
        }
        actorRows.Close() 

        movie.Actors = actors 
        movies = append(movies, movie)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return movies, nil
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
func (db *API) AddMovie(ctx context.Context, movie models.MovieResponse) (models.MovieResponseActor, error) {
	var existingMovieID int
	movieRating, err := strconv.ParseFloat(movie.Rating, 3)
	truncatedStr := fmt.Sprintf("%.1f", movieRating)
	movieRating, err = strconv.ParseFloat(truncatedStr, 3)

	checkQuery := `SELECT id FROM movies WHERE title = $1 AND release_date = $2 AND rating = $3`
    err = db.db.QueryRowContext(ctx, checkQuery, movie.Title, movie.ReleaseDate, movieRating).Scan(&existingMovieID)
    if err != nil {
    } else {
        return models.MovieResponseActor{}, fmt.Errorf("The film already exists")
    }
    for _, actorIDStr := range movie.Actors {
        var actorID int
        if _, err := fmt.Sscanf(actorIDStr, "%d", &actorID); err != nil {
            return models.MovieResponseActor{}, fmt.Errorf("invalid actor ID format: %v", err)
        }

        var exists bool
        err = db.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM actors WHERE id = $1)", actorID).Scan(&exists)
        if err != nil || !exists {
            return models.MovieResponseActor{}, fmt.Errorf("actor with ID %d not found", actorID)
        }
    }

	var movieID int
	query := `
		INSERT INTO movies (title, description, release_date, rating)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	err = db.db.QueryRowContext(ctx, query, movie.Title, movie.Description, movie.ReleaseDate, movieRating).Scan(&movieID)
	if err != nil {
		return models.MovieResponseActor{}, err
	}

    for _, actorIDStr := range movie.Actors {
        var actorID int
        if _, err := fmt.Sscanf(actorIDStr, "%d", &actorID); err != nil {
            return models.MovieResponseActor{}, fmt.Errorf("invalid actor ID format: %v", err)
        }

        var exists bool
        err = db.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM actors WHERE id = $1)", actorID).Scan(&exists)
        if err != nil || !exists {
            return models.MovieResponseActor{}, fmt.Errorf("actor with ID %d not found", actorID)
        }

        _, err = db.db.ExecContext(ctx, "INSERT INTO actors_movies (actor_id, movie_id) VALUES ($1, $2)", actorID, movieID)
        if err != nil {

            return models.MovieResponseActor{}, fmt.Errorf("failed to insert actor-movie relation: %v", err)
        }
    }
	var actors []models.Actor
	for _, actorIDStr := range movie.Actors {
	var actor models.Actor
	actorQuery := `SELECT * FROM actors WHERE id = $1;`
        err := db.db.QueryRowContext(ctx, actorQuery, actorIDStr).Scan(&actor.Id, &actor.Name, &actor.BirthDate, &actor.BirthDate)
        if err != nil {
            if err == sql.ErrNoRows {
                continue // Актёр не найден, переходим к следующему ID
            }
            return models.MovieResponseActor{}, fmt.Errorf("querying actor info: %v", err)
        }
	actors = append(actors, actor)
	}


    // Возвращаем добавленный фильм
    return models.MovieResponseActor{
        Id:          fmt.Sprintf("%d", movieID),
        Title:       movie.Title,
        Description: movie.Description,
        ReleaseDate: movie.ReleaseDate,
        Rating:      movie.Rating,
		Actors: 	actors,
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
	checkQuery := `SELECT id FROM users WHERE username = $1`
    var existingUserId string
    err := db.db.QueryRowContext(ctx, checkQuery, user.Username).Scan(&existingUserId)
    if err == nil {
        return models.UserResponse{}, fmt.Errorf("user with username %s already exists", user.Username)
    }

	query := `
		INSERT INTO users (username, refresh_token, role)
		VALUES ($1, $2, $3)
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
	err = db.db.QueryRowContext(ctx, query, user.Username, string(refreshToken), user.Role).Scan(&userID)
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


