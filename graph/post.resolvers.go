package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ipreferwater/doko-graphql/db"
	"github.com/ipreferwater/doko-graphql/model"
)

func (r *mutationResolver) CreatePosts(ctx context.Context, input []*model.InputPost) (string, error) {
	db.PostRepository.CreatePosts(input)
	return "ok", nil
}

func (r *queryResolver) Posts(ctx context.Context) (*model.GetPosts, error) {
	posts, err := db.PostRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	return &model.GetPosts{
		Posts: posts,
	}, nil
}
