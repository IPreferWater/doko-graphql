package db

import (
	"github.com/ipreferwater/graphql-theory/graph/model"
)

type NoteRepositoryInterface interface {
	GetNotes() ([]*model.Note, error)
}
