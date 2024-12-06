package domains

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"song-library/models"
)

const (
	songTable = "songs"
)

type Repository interface {
	Create(song models.Song) (int, error)
	Update(id int, song models.Song) (int, error)
	Delete(id int) error
	GetSliceSongs(offset, limit int) ([]models.Song, error)
	GetSongs() ([]models.Song, error)
	GetLyrics(songId int) (string, error)
}

type SongRepository struct {
	db *sqlx.DB
	Repository
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(song models.Song) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (author, title, release_date, lyrics, link) VALUES ($1, $2, $3, $4, $5) RETURNING id", songTable)
	row := r.db.QueryRow(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link)

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
func (r *SongRepository) Update(id int, song models.Song) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET author = $1, title = $2, release_date = $3, lyrics = $4, link = $5 WHERE id = $6", songTable)

	_, err := r.db.Exec(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link, id)
	if err != nil {
		return -1, err
	}
	return song.Id, nil
}

func (r *SongRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", songTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SongRepository) GetSongs() ([]models.Song, error) {
	var songs []models.Song
	query := fmt.Sprintf("SELECT * FROM %s", songTable)
	err := r.db.Select(&songs, query)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongRepository) GetSliceSongs(offset, limit int) ([]models.Song, error) {
	var songs []models.Song
	query := fmt.Sprintf("SELECT * FROM %s WHERE id >= $1 and id < ($1 + $2)", songTable)
	err := r.db.Select(&songs, query, offset, limit)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongRepository) GetLyrics(songId int) (string, error) {
	query := fmt.Sprintf("SELECT text FROM %s WHERE id=$1", songTable)
	var lyrics string
	err := r.db.Get(&lyrics, query, songId)
	if err != nil {
		return "", err
	}
	return lyrics, nil
}
