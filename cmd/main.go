package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/configs"
	"song-library/db"
	"song-library/internal/domains"
	"song-library/internal/handlers"
	"song-library/internal/services"
	"song-library/logger"
)

func main() {
	var err error
	logger.Log, err = logger.GetNewLogger()

	if err != nil {
		panic(err)
	}
	logger.Log.Info("Init logger successful")

	database, err := db.NewDB(configs.GetConfig())
	if err != nil {
		logger.Log.Fatal("Init database was failed", err)
		panic(err)
	}
	err = database.Ping()
	if err != nil {
		logger.Log.Fatal("Init database was failed", err)
	}
	logger.Log.Info("Init database successful")

	//Internal layer - внутренний слой
	repository := domains.NewSongRepository(database)
	service := services.NewSongService(repository)
	handler := handlers.NewHandler(service)
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
}
