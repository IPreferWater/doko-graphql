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
	/*var arrNotes []*model.Note
	for i := 0; i < 10; i++ {

		var arrSteps []model.Step
		for j := 0; j < 5; j++ {
			step := model.Step{
				Title: fmt.Sprintf("title-%d", j),
				Txt:   fmt.Sprintf("text-%d", j),
			}

			if j == 3 {
				urlString := fmt.Sprintf("url-%d", j)
				step.Url = &urlString
			}

			arrSteps = append(arrSteps, step)
		}
		arrNotes = append(arrNotes, &model.Note{
			Name:  fmt.Sprintf("notes-%d", i),
			Steps: arrSteps,
		})
	}
	return arrNotes, nil*/
}

func (n TODONoteRepository) CreateNote(newNote model.Note) error {

	allNotes = append(allNotes, &newNote)
	return nil
}

func InitTODORepo() {

	NoteRepository = &TODONoteRepository{}
	//log.Info("connected")
}
