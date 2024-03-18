package queries

import (
	"context"
	"movie_storage/internal/models"
	. "movie_storage/internal/models"
)



type Storage interface {
	GetMovies(ctx context.Context, sortParam, orderParam, titleParam, actorNameParam string) ([]Movie, error)
	AddMovie(ctx context.Context, movie models.MovieResponse) (Movie, error)
	DeleteMovie(ctx context.Context, movieID string) error
	UpdateMovie(ctx context.Context, movieID string, movie models.MovieResponse) (Movie, error)
	GetActorMovies(ctx context.Context, actorID string) ([]Movie, error)
	AddActor(ctx context.Context, actor ActorResponse) (Actor, error)
	DeleteActor(ctx context.Context, actorID string) error
	UpdateActor(ctx context.Context, actorID string, actor ActorResponse) (Actor, error)
	RegisterUser(ctx context.Context, user RegisterRequest) (models.UserResponse, error)
	LoginUser(ctx context.Context, userId string, user LoginRequest) (models.UserResponse, error)
	GetToken(ctx context.Context) (string, error)
	RefreshToken(ctx context.Context, refresh string) (string, error)
 }