package domains

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"song-library/models"
)

const (
	songTable = "songs"
)

type Song interface {
	Create(song models.Song) (uuid.UUID, error)
	Update(c *gin.Context) (uuid.UUID, error)
	Delete(c *gin.Context) error
	GetSliceSongs(c *gin.Context, offset, limit int) []models.Song
	GetSongs(c *gin.Context) []models.Song
}

type SongRepository struct {
	db *sqlx.DB
	Song
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(song models.Song) (uuid.UUID, error) {
	query := fmt.Sprintf("INSERT INTO %s (author, title, release_date, lyrics, link) VALUES ($1, $2, $3, $4, $5) RETURNING id", songTable)
	row := r.db.QueryRow(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
func (r *SongRepository) Update(song models.Song) (uuid.UUID, error) {
	query := fmt.Sprintf("UPDATE %s SET author = $1, title = $2, release_date = $3, lyrics = $4, link = $5 WHERE id = $6", songTable)

	_, err := r.db.Exec(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link, song.Id)
	if err != nil {
		return uuid.Nil, err
	}
	return song.Id, nil
}

func (r *SongRepository) Delete(c *gin.Context) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", songTable)
	_, err := r.db.Exec(query, c.Param("id"))
	if err != nil {
		return err
	}
	return nil
}
func (r *SongRepository) GetSongs(c *gin.Context) ([]models.Song, error) {
	var songs []models.Song
	query := fmt.Sprintf("SELECT * FROM %s", songTable)
	err := r.db.Select(&songs, query)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
