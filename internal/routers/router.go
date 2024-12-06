package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/internal/handlers"
	"song-library/logger"
)

func NewRouter(handler handlers.Handler) *gin.Engine {
	router := gin.Default()

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

	logger.Log.Info("Init router successful")
	return router
}
