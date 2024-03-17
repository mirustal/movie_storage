package queries

import (
	"context"
	. "movie_storage/internal/models"
)



type Storage interface {
	GetMovies(ctx context.Context, sortParam, orderParam, titleParam, actorNameParam string) ([]Movie, error)
	AddMovie(ctx context.Context, movie Movie) (Movie, error)
	DeleteMovie(ctx context.Context, movieID string) error
	UpdateMovie(ctx context.Context, movieID string, movie Movie) (Movie, error)
	GetActorMovies(ctx context.Context, actorID string) ([]Movie, error)
	AddActor(ctx context.Context, actor Actor) (Actor, error)
	DeleteActor(ctx context.Context, actorID string) error
	UpdateActor(ctx context.Context, actorID string, actor Actor) (Actor, error)
	RegisterUser(ctx context.Context, user RegisterRequest) (string, error)
	GetToken(ctx context.Context) (string, error)
	RefreshToken(ctx context.Context, refresh string) (string, error)
 }