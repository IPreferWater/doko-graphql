package db

import (
	"github.com/ipreferwater/graphql-theory/model"
)

type NoteRepositoryInterface interface {
	GetNotes() ([]*model.Note, error)
	CreateNote(model.Note) error
}
