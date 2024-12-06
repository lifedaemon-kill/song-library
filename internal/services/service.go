package services

import (
	"song-library/internal/domains"
	"song-library/models"
	"strings"
)

type Service interface {
	DeleteSong(id int) error
	GetLyrics(songId, pageOffset int) (string, error)
	UpdateSong(id int, song models.Song) (int, error)
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

	return nil
}

func (s *SongService) GetLyrics(songId, pageOffset int) (string, error) {
	const sizeVerse = 3 //Как много куплетов выводить
	lyrics, err := s.repo.GetLyrics(songId)
	if err != nil {
		return "", err
	}

	verses := strings.Split(lyrics, "\n\n")
	targetVerses := verses[pageOffset : pageOffset+sizeVerse]

	return strings.Join(targetVerses, "\n\n"), nil
}
func (s *SongService) UpdateSong(id int, song models.Song) (int, error) {
	updatedId, err := s.repo.Update(id, song)
	if err != nil {
		return -1, err
	}
	return updatedId, nil
}
