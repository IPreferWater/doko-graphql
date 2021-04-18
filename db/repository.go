package db

import (
	"github.com/ipreferwater/doko-graphql/model"
)


var (
	PostRepository PostRepositoryInterface
)

type NoteRepositoryInterface interface {
	GetNotes() ([]*model.Note, error)
	CreateNote(model.Note) error
	UpdateNote(model.Note) error
	DeleteNote(int) error
}

type PostRepositoryInterface interface {
	GetPosts() ([]model.Post, error)
	CreatePosts([]*model.InputPost) error
	UpdatePosts([]model.Post) error
	DeletePosts([]int) error

	GetUserIdByUsernamePassword(userName string, password string) (int,error)
}
