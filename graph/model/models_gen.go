// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewNote struct {
	Name  string     `json:"name"`
	Steps []*NewStep `json:"steps"`
}

type NewStep struct {
	Title string  `json:"title"`
	Txt   string  `json:"txt"`
	URL   *string `json:"url"`
}