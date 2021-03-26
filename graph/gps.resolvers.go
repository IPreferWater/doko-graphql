package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ipreferwater/graphql-theory/auth"
	"github.com/ipreferwater/graphql-theory/db"
	model1 "github.com/ipreferwater/graphql-theory/graph/model"
	"github.com/ipreferwater/graphql-theory/model"
)

func (r *mutationResolver) CreatePosts(ctx context.Context, input []*model.InputPost) (string, error) {
	db.PostRepository.CreatePosts(input)
	return "ok", nil
}

func (r *mutationResolver) Login(ctx context.Context, input model1.InputLogin) (string, error) {
	id, err := db.PostRepository.GetUserIdByUsernamePassword(input.Username, input.Password)
	if err != nil {
		return "error find user", err
	}

	if id < 1 {
		return "not found", err
	}

	token, err := auth.CreateToken(id)
	if err != nil {
		return "error create token", err
	}

	return token, err

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
