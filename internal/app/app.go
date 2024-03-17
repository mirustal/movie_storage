package app

import (
	"log"
	"movie_storage/internal/handler"
	"movie_storage/internal/router"
	"movie_storage/internal/validation"
	"movie_storage/pkg/configs"
	"movie_storage/platform/database"
)

type App struct {
	Router *router.Router
	Config *configs.Config
}


func NewApp(cfg *configs.Config) *App {
	db, err := database.NewDatabase(&cfg.PostgresDB)
	if err != nil {
		return &App{}
	}

	storage := database.NewAPI(db.GetDB())
	validation := validation.NewValidator(storage)
	handlers := handler.NewHandler(validation)

	router := router.NewRouter(cfg, handlers)

	return &App{
		Router: router,
		Config: cfg,
	}
}

func (app *App) Run() {
	log.Printf("server running %s", app.Config.Http.Port)
	if err := app.Router.Run(app.Config.Http.Port); err != nil {
		log.Fatal(err)
	}
}
