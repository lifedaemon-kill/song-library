package services

import (
	"encoding/json"
	"io"
	"net/http"
	"song-library/models"
)

type ClientService interface {
	GetDetailData(params models.InfoQueryParams) (models.SongDetail, error)
}

type clientService struct {
	apiHost string //swagger host path /info
}

func NewClientService(host string) ClientService {
	return &clientService{host}
}

func (c *clientService) GetDetailData(params models.InfoQueryParams) (models.SongDetail, error) {
	response, err := http.Get(c.apiHost + "/info")
	if err != nil {
		return models.SongDetail{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.SongDetail{}, err
	}

	var details models.SongDetail
	if err = json.Unmarshal(body, &details); err != nil {
		return models.SongDetail{}, err
	}

	return details, nil
}
