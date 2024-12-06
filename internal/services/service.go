package services

import (
	"song-library/internal/domains"
	"strings"
)

type Service interface {
	DeleteSong(id int) error
	GetLyrics(songId, pageOffset int) (string, error)
}

type SongService struct {
	repo domains.Repository
}

func NewSongService(repo domains.Repository) *SongService {
	return &SongService{repo}
}

func (s *SongService) DeleteSong(id int) error {
	err := s.repo.Delete(id)

	if err != nil {
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
