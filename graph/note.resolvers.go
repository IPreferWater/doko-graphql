package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ipreferwater/graphql-theory/db"
	"github.com/ipreferwater/graphql-theory/graph/generated"
	"github.com/ipreferwater/graphql-theory/graph/model"
	model1 "github.com/ipreferwater/graphql-theory/model"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (string, error) {

	var newSteps []model1.Step
	for _, newStep := range input.Steps {
		mapedStep := model1.Step{
			Title: newStep.Title,
			Txt:   newStep.Txt,
		}

		if newStep.URL != nil {
			mapedStep.Url = newStep.URL
		}
		newSteps = append(newSteps, mapedStep)

	}
	newNote := model1.Note{
		Name:  input.Name,
		Steps: newSteps,
	}
	if err := db.NoteRepository.CreateNote(newNote); err != nil {
		return "error", nil
	}

	return "done", nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]*model1.Note, error) {
	log.Info("query notes")
	return db.NoteRepository.GetNotes()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }