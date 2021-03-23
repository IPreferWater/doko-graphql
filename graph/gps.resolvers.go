package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ipreferwater/graphql-theory/db"
	model1 "github.com/ipreferwater/graphql-theory/graph/model"
	"github.com/ipreferwater/graphql-theory/model"
)

func (r *mutationResolver) CreatePosts(ctx context.Context, input model.InputPost) (string, error) {
	
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdatePosts(ctx context.Context, input model1.InputUpdatePosts) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DeletePostsByID(ctx context.Context, input []int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
