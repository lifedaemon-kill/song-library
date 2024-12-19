package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/pressly/goose"
	"net/http"
	"os"
	"os/signal"
	"song-library/configs"
	"song-library/internal/domains"
	"song-library/internal/handlers"
	"song-library/internal/pkg/logger"
	"song-library/internal/pkg/storage"
	"song-library/internal/routers"
	"song-library/internal/services"
	"sync"
	"syscall"
	"time"
)

const migrationsPath = "db/migrations"

func main() {
	var err error
	//Logger
	logger.Log, err = logger.GetNewLogger()
	if err != nil {
		fmt.Println(err)
		return
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
	clientService := services.NewClientService(config.Server.ExternalApiAddr)
	handler := handlers.NewHandler(service, clientService)

	logger.Log.Info("Init internal layer successful")

	//Router
	router := routers.NewRouter(handler)

	server := &http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}
	//Start server
	go func() {
		logger.Log.Info("Starting Gin server at localhost:8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Log.Fatalf("Failed to listen and serve: %v", err)
		}
	}()
	logger.Log.Info("Server start completed")

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdown()

		if err := server.Shutdown(ctx); err != nil {
			logger.Log.Error("failed to stop server: %v", err)
		}

		if err := database.Close(); err != nil {
			logger.Log.Error("failed to close database: %v", err)
		}

		logger.Log.Info("Graceful shutdown complete")
	}()

	wg.Wait()

}
