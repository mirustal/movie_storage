package main

import (
	"movie_storage/internal/app"
	"movie_storage/pkg/configs"
	"os"
)

func main() {

	os.Setenv("CONFIG_PATH", "/Users/mirustal/Documents/project/go/movie_storage/config.yml")
	os.Setenv("SECRET_KEY", "Medods_Task1")

	cfg := configs.GetConfig()
	// log := logging.SetupLogger(cfg.ModeLog)
	// log.Info("Starting service", slog.String("env", cfg.ModeLog))
	app := app.NewApp(cfg)
	app.Run()

}

