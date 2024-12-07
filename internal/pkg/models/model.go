package models

type Song struct {
	Id     int    `json:"id" db:"id"`
	Author string `json:"group" db:"author"`
	Title  string `json:"song" db:"title"`
	SongDetail
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate" db:"release_date"` //example: 16.07.2006
	Lyrics      string `json:"text" db:"lyrics"`              //example: Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses
	Link        string `json:"link" db:"link"`                //example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
}

type InfoQueryParams struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type FilterParams struct {
	ID          *int    `form:"id"`
	Author      *string `form:"group"`
	Title       *string `form:"title"`
	ReleaseDate *string `form:"release_date"`
	Lyrics      *string `form:"lyrics"`
	Link        *string `form:"link"`
}
