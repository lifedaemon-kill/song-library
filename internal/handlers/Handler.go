package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/internal/services"
	"song-library/logger"
	"strconv"
)

type Handler interface {
	GetLibrary(c *gin.Context)
	GetLyrics(c *gin.Context)
	DeleteSong(c *gin.Context)
	UpdateSong(c *gin.Context)
	AddSong(c *gin.Context)
}
type SongHandler struct {
	songService services.Service
}

func NewHandler(songService services.Service) *SongHandler {
	return &SongHandler{songService}
}

// GetLibrary Получение данных библиотеки с фильтрацией по всем полям и пагинацией
func (h *SongHandler) GetLibrary(c *gin.Context) {

}

// GetLyrics Получение текста песни с пагинацией по куплетам
func (h *SongHandler) GetLyrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offsetRow := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetRow)
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lyrics, err := h.songService.GetLyrics(id, offset)
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, lyrics)
}

// DeleteSong Удаление песни
func (h *SongHandler) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.songService.DeleteSong(id)

	if err != nil {
		logger.Log.Error(err)
		c.JSON(500, gin.H{"error": err})
		return
	}
	logger.Log.Info("Song " + strconv.Itoa(id) + " deleted")
	c.JSON(200, gin.H{"message": "ok"})
}
