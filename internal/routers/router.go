package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"song-library/internal/handlers"
	"song-library/logger"
)

func NewRouter(handler handlers.Handler) *gin.Engine {
	router := gin.Default()

	// Маршрут для статического файла OpenAPI
	router.StaticFile("/openapi.yaml", "./openapi.yaml")

	// Маршрут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/openapi.yaml")))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
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
