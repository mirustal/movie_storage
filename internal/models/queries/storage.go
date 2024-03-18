package queries

import (
	"context"
	"movie_storage/internal/models"
	. "movie_storage/internal/models"
)



type Storage interface {
	GetMovies(ctx context.Context, typeGet string) ([]MovieResponseActor, error)
	SearchMovies(ctx context.Context, search string) ([]MovieResponseActor, error)
	AddMovie(ctx context.Context, movie models.MovieResponse) (MovieResponseActor, error)
	DeleteMovie(ctx context.Context, movieID string) error
	UpdateMovie(ctx context.Context, movieID string, movie models.MovieResponse) (Movie, error)
	GetActorsMovies(ctx context.Context, actorID models.ActorsIdRequest) ([]ActorWithMovies, error)
	AddActor(ctx context.Context, actor ActorResponse) (Actor, error)
	DeleteActor(ctx context.Context, actorID string) error
	UpdateActor(ctx context.Context, actorID string, actor ActorResponse) (Actor, error)
	RegisterUser(ctx context.Context, user RegisterRequest) (models.UserResponse, error)
	LoginUser(ctx context.Context, userId string, user LoginRequest) (models.UserResponse, error)
	GetToken(ctx context.Context) (string, error)
	RefreshToken(ctx context.Context, refresh string) (string, error)
 }