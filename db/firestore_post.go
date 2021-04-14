package db

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/ipreferwater/doko-graphql/model"
)

type FirestorePostRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func (f FirestorePostRepository) GetPosts() ([]model.Post, error) {
	iter := f.client.Collection("posts").Documents(f.ctx)
	defer iter.Stop()

	docs, err := iter.GetAll()
	if err != nil {
		return nil, err
	}

	var res []model.Post
	for _, doc := range docs {
		var post model.Post
		if err := doc.DataTo(&post); err != nil {
			return nil, err
		}
		res = append(res, post)
	}

	return res, nil
}

func (f FirestorePostRepository) CreatePosts(newPosts []*model.InputPost) error {
	collection := f.client.Collection("posts")

	writeBatch := f.client.Batch()
	for _, newPost := range newPosts {
		now := time.Now()
		docRef := collection.Doc(now.String())
		writeBatch.Create(docRef, *newPost)
	}

	_, err := writeBatch.Commit(f.ctx)
	return err
}

func (f FirestorePostRepository) UpdatePosts(postsToUpdate []model.Post) error {
	return nil
}
func (f FirestorePostRepository) DeletePosts(idsPostToDelete []int) error {
	return nil
}

func (f FirestorePostRepository) GetUserIdByUsernamePassword(userName string, password string) (int, error) {
	var id int
	id = 1
	return id, nil
}

func InitFirestorePostRepository() {
	ctx := context.TODO()

	client, err := firestore.NewClient(ctx, "dummy-project-id")
	if err != nil {
		panic(err)
	}

	PostRepository = &FirestorePostRepository{client: client, ctx: ctx}
}
