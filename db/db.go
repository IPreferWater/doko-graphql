package db

import (
	//"github.com/ipreferwater/graphql-theory/model"
	"github.com/ipreferwater/graphql-theory/model"
)

var (
	NoteRepository NoteRepositoryInterface

	allNotes []*model.Note
)

type TODONoteRepository struct {
}

func (n TODONoteRepository) GetNotes() ([]*model.Note, error) {

	return allNotes, nil
}

func (n TODONoteRepository) CreateNote(newNote model.Note) error {

	allNotes = append(allNotes, &newNote)
	return nil
}

func (n TODONoteRepository) UpdateNote(newNote model.Note) error {

	allNotes[len(allNotes)] = &newNote
	return nil
}

func (n TODONoteRepository) DeleteNote(id int) error {

	if len(allNotes) >= 1 {
		allNotes = allNotes[:len(allNotes)-1]
	}
	return nil
}

func InitTODORepo() {

	NoteRepository = &TODONoteRepository{}
	//log.Info("connected")
}
