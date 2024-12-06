package services

import (
	"net/http"
	"song-library/models"
)

type ClientService interface {
	GetDetailData(params models.InfoQueryParams) (models.Song, error)
}

type clientService struct {
	apiHost string //swagger host path /info
}

func NewClientService(host string) ClientService {
	return &clientService{host}
}

func (c *clientService) GetDetailData(params models.InfoQueryParams) (models.Song, error) {
	response, err := http.Get(c.apiHost + "/info")
	if err != nil {
		return models.Song{}, err
	}
	defer response.Body.Close()
	var ResponseSongDetail models.SongDetail

	song := models.Song{
		Author:     params.Group,
		Title:      params.Song,
		SongDetail: ResponseSongDetail,
	}

	return song, nil
}
