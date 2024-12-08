package domains

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"song-library/internal/pkg/logger"
	"song-library/internal/pkg/models"
	"strings"
)

const (
	songTable = "songs"
)

type Repository interface {
	Create(song models.Song) (int, error)
	Update(id int, song models.Song) (int, error)
	Delete(id int) error
	GetLyrics(songId int) (string, error)
	GetFilteredLib(params models.FilterParams) ([]models.Song, error)
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
	logger.Log.Debug("Create query=", query, "song=", song)

	row := r.db.QueryRow(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link)

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
func (r *SongRepository) Update(id int, song models.Song) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET author = $1, title = $2, release_date = $3, lyrics = $4, link = $5 WHERE id = $6", songTable)

	logger.Log.Debug("Delete query=", query, "song=", song)

	if _, err := r.db.Exec(query, song.Author, song.Title, song.ReleaseDate, song.Lyrics, song.Link, id); err != nil {
		return -1, err
	}
	return song.Id, nil
}

func (r *SongRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", songTable)

	logger.Log.Debug("Delete query=", query)

	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (r *SongRepository) GetLyrics(songId int) (string, error) {
	query := fmt.Sprintf("SELECT lyrics FROM %s WHERE id=$1", songTable)
	var lyrics string

	logger.Log.Debug("Get Lyrics query=", query)

	if err := r.db.Get(&lyrics, query, songId); err != nil {
		return "", err
	}
	return lyrics, nil
}

func (r *SongRepository) GetFilteredLib(params models.FilterParams) ([]models.Song, error) {
	query := "SELECT id, author, title, release_date, lyrics, link FROM " + songTable + " WHERE 1=1"
	var args []interface{}

	if params.ID != nil {
		query += " AND id = ?"
		args = append(args, *params.ID)
	}

	if params.Author != nil {
		query += " AND author LIKE ?"
		args = append(args, "%"+*params.Author+"%")
	}

	if params.Title != nil {
		query += " AND title LIKE ?"
		args = append(args, "%"+*params.Title+"%")
	}

	if params.ReleaseDate != nil {
		query += " AND release_date = ?"
		args = append(args, *params.ReleaseDate)
	}

	if params.Lyrics != nil {
		query += " AND lyrics LIKE ?"
		args = append(args, "%"+*params.Lyrics+"%")
	}

	if params.Link != nil {
		query += " AND link LIKE ?"
		args = append(args, "%"+*params.Link+"%")
	}

	query += " ORDER BY id ASC"

	if params.Limit != nil {
		query += " LIMIT ?"
		args = append(args, *params.Limit)
	} else {
		query += " LIMIT ?"
		args = append(args, 10)
	}

	if params.Offset != nil {
		query += " OFFSET ?"
		args = append(args, *params.Offset)
	} else {
		query += " OFFSET ?"
		args = append(args, 0)
	}

	var songs []models.Song
	logger.Log.Debug("Get Lib query=", query, " args=", args)
	finalQuery := query
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			finalQuery = strings.Replace(finalQuery, "?", fmt.Sprintf("'%s'", v), 1)
		default:
			finalQuery = strings.Replace(finalQuery, "?", fmt.Sprintf("%v", v), 1)
		}
	}
	logger.Log.Debug("Get Lib FINAL query=", finalQuery)

	if err := r.db.Select(&songs, finalQuery); err != nil {
		return nil, err
	}

	return songs, nil
}
