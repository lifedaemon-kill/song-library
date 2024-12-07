package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"song-library/internal/pkg/logger"
	"song-library/internal/pkg/models"
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
	requestURL := fmt.Sprintf("http://%s/info?group=%s&song=", c.apiHost, params.Group, params.Song)
	response, err := http.Get(requestURL)

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

	logger.Log.Debug("info about ", params, " received")
	return details, nil
}
