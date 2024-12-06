package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/internal/services"
	"song-library/logger"
	"song-library/models"
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
	songService   services.Service
	clientService services.ClientService
}

func NewHandler(songService services.Service, clientService services.ClientService) *SongHandler {
	return &SongHandler{songService, clientService}
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

	if err = h.songService.DeleteSong(id); err != nil {
		logger.Log.Error(err)
		c.JSON(500, gin.H{"error": err})
		return
	}
	logger.Log.Info("Song " + strconv.Itoa(id) + " deleted")
	c.JSON(200, gin.H{"message": "ok"})
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var song models.Song
	if err = c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id, err = h.songService.UpdateSong(id, song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
func (h *SongHandler) AddSong(c *gin.Context) {
	var info models.InfoQueryParams
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	details, err := h.clientService.GetDetailData(info)

	song := models.Song{
		Title:      info.Song,
		Author:     info.Group,
		SongDetail: details,
	}

	var id int
	id, err = h.songService.AddSong(song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
