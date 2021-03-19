package db

import (
	"fmt"

	"github.com/ipreferwater/graphql-theory/graph/model"
)

var (
	NoteRepository NoteRepositoryInterface
)

type TODONoteRepository struct {
}

func (n TODONoteRepository) GetNotes() ([]*model.Note, error) {

	var arrNotes []*model.Note
	for i := 0; i < 10; i++ {

		var arrSteps []*model.Step
		for j := 0; j < 5; j++ {
			step := model.Step{
				Title: fmt.Sprintf("title-%d", j),
				Txt:   fmt.Sprintf("text-%d", j),
			}

			if j == 3 {
				urlString := fmt.Sprintf("url-%d", j)
				step.URL = &urlString
			}

			arrSteps = append(arrSteps, &step)
		}
		arrNotes = append(arrNotes, &model.Note{
			Name:  fmt.Sprintf("notes-%d", i),
			Steps: arrSteps,
		})
	}
	return arrNotes, nil
}

func InitTODORepo() {

	NoteRepository = &TODONoteRepository{}
	//log.Info("connected")
}
