package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
	"net/http"
	"song-library/configs"
	"song-library/db"
	"song-library/internal/domains"
	"song-library/internal/handlers"
	"song-library/internal/services"
	"song-library/logger"
)

const migrationsPath = "migrations"

func main() {
	var err error
	logger.Log, err = logger.GetNewLogger()

	if err != nil {
		panic(err)
	}
	logger.Log.Info("Trying start server...")

	database, err := db.NewDB(configs.GetConfig())
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

	//Internal layer - внутренний слой
	repository := domains.NewSongRepository(database)
	service := services.NewSongService(repository)
	clientService := services.NewClientService("localhost:3000")
	handler := handlers.NewHandler(service, clientService)
	logger.Log.Info("Init internal layer successful")

	//Server
	router := gin.Default()
	logger.Log.Info("Init server successful")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong!")
	})

	//1 Получение данных библиотеки с фильтрацией по всем полям и пагинацией
	router.GET("/songs", handler.GetLibrary)

	//2 Получение текста песни с пагинацией по куплетам
	router.GET("/songs/:id", handler.GetLyrics)

	//3 Удаление песни
	router.DELETE("songs/:id", handler.DeleteSong)

	//4 Изменение данных песни
	router.PUT("songs/:id", handler.UpdateSong)

	//5 Добавление новой песни в формате JSON
	router.POST("/songs", handler.AddSong)

	logger.Log.Info("Server start completed")
}
