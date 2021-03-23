package model

type GetPosts struct {
	Posts    []Post     `json:"posts"`
}

type NewPosts struct {
	NewPosts []InputPost `json:"newPosts"`
}

type InputPost struct {
	Title string    `json:"title"`
	Txt   *string   `json:"txt"`
	Gps   Gps `json:"gps"`
}

type InputGps struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}



type Post struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Txt   *string `json:"txt"`
	Gps   Gps     `json:"gps"`
}

type Gps struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
