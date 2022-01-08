package models

type Category struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
}

type Film struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	CategoryID int32  `json:"category_id"`
}

type FilmCategory struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Category   string `json:"category"`
}
