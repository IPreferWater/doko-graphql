package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ipreferwater/graphql-theory/db"
	"github.com/ipreferwater/graphql-theory/graph/generated"
	model1 "github.com/ipreferwater/graphql-theory/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Notes(ctx context.Context) ([]*model1.Note, error) {
	log.Info("query notes")
	return db.NoteRepository.GetNotes()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
