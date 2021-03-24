package db

import (
	"github.com/ipreferwater/graphql-theory/model"
)

var (
	PostRepository PostRepositoryInterface
	allPosts       []model.Post
)

type MysqlPostRepository struct {
}

func (n MysqlPostRepository) GetPosts() ([]model.Post, error) {
	return allPosts, nil
}

func (n MysqlPostRepository) CreatePosts(newPosts []*model.InputPost) error {

	for _, newPost := range newPosts {
		id := len(allPosts)
		allPosts = append(allPosts, model.Post{
			ID:    id,
			Title: newPost.Title,
			Txt:   newPost.Txt,
			Gps: model.Gps{
				X: newPost.Gps.X,
				Y: newPost.Gps.Y,
			},
		})
	}
	return nil
}

func (n MysqlPostRepository) UpdatePosts(postsToUpdate []model.Post) error {
	return nil
}
func (n MysqlPostRepository) DeletePosts(idsPostToDelete []int) error {
	return nil
}

func InitMysqlPostRepository() {
	PostRepository = &MysqlPostRepository{}
}
