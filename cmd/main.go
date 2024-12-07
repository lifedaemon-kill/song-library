package main

import (
	"github.com/pressly/goose"
	"song-library/configs"
	"song-library/internal/domains"
	"song-library/internal/handlers"
	"song-library/internal/pkg/logger"
	"song-library/internal/pkg/storage"
	"song-library/internal/routers"
	"song-library/internal/services"
)

const migrationsPath = "db/migrations"

func main() {
	var err error
	//Logger
	logger.Log, err = logger.GetNewLogger()
	if err != nil {
		panic(err)
	}
	logger.Log.Info("Trying start server...")

	config := configs.GetConfig()

	//Database
	database, err := storage.NewDB(config)
	if err != nil {
		logger.Log.Fatal("Init database was failed | ", err)
		panic(err)
	}
	logger.Log.Info("Init database successful")

	//migrations
	if err := goose.Up(database.DB, migrationsPath); err != nil {
		logger.Log.Info("Init migrations was failed")
		panic(err)
	}
	logger.Log.Info("Init migrations successful")

	//Internal layer
	repository := domains.NewSongRepository(database)
	service := services.NewSongService(repository)
	clientService := services.NewClientService(config.ExternalApi)
	handler := handlers.NewHandler(service, clientService)

	logger.Log.Info("Init internal layer successful")

	//Router
	router := routers.NewRouter(handler)

	if err := router.Run("localhost:8080"); err != nil {
		logger.Log.Fatal("router run failed")
		return
	}
	logger.Log.Info("Server start completed")
}
