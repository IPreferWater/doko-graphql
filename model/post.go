package model

type GetPosts struct {
	Posts []Post `json:"posts"`
}

type NewPosts struct {
	NewPosts []InputPost `json:"newPosts"`
}

type InputPost struct {
	Title     string  `json:"title"`
	Text       *string `json:"txt"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Post struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Text       *string `json:"txt"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
