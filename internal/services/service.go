package services

import (
	"errors"
	"song-library/internal/domains"
	"song-library/internal/pkg/logger"
	"song-library/internal/pkg/models"
	"strings"
)

type Service interface {
	DeleteSong(id int) error
	GetLyrics(songId, pageOffset int) (string, error)
	UpdateSong(id int, song models.Song) (int, error)
	AddSong(song models.Song) (int, error)
	GetFilteredLib(params models.FilterParams) ([]models.Song, error)
}

type SongService struct {
	repo domains.Repository
}

func NewSongService(repo domains.Repository) *SongService {
	return &SongService{repo}
}

func (s *SongService) DeleteSong(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	logger.Log.Debug("song ", id, " deleted")
	return nil
}

func (s *SongService) GetLyrics(songId, from int) (string, error) {
	const verseCount = 3 //Сколько куплетов выводить
	lyrics, err := s.repo.GetLyrics(songId)
	if err != nil {
		return "", err
	}

	verses := strings.Split(lyrics, "\n\n")
	if len(verses) < from {
		return "", errors.New("invalid offset")
	}
	var to int
	if len(verses) == 0 {

	}
	if from+verseCount > len(verses) {
		to = len(verses)
	} else {
		to = from + verseCount
	}

	targetVerses := verses[from:to]
	logger.Log.Debug("get lyr=", from, to, targetVerses, lyrics)

	logger.Log.Debug("lyrics for song ", songId, " recieved")
	return strings.Join(targetVerses, "\n\n"), nil
}

func (s *SongService) UpdateSong(id int, song models.Song) (int, error) {
	updatedId, err := s.repo.Update(id, song)
	if err != nil {
		return -1, err
	}
	logger.Log.Debug("song updated with id: %d", updatedId)
	return updatedId, nil
}
func (s *SongService) AddSong(song models.Song) (int, error) {
	id, err := s.repo.Create(song)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("song added with id: ", id)
	return id, nil
}
func (s *SongService) GetFilteredLib(params models.FilterParams) ([]models.Song, error) {
	lib, err := s.repo.GetFilteredLib(params)
	if err != nil {
		return nil, err
	}
	logger.Log.Debug("filtered songs received. params:", params)
	return lib, nil
}
