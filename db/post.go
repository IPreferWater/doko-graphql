package db

import "github.com/ipreferwater/graphql-theory/model"

var (
	PostRepository PostRepositoryInterface
	allPosts       []model.Post
)

type MysqlPostRepository struct {
}

func (n MysqlPostRepository) GetPosts() ([]model.Post, error) {
	return allPosts, nil
}

func (n MysqlPostRepository) CreatePosts(newPosts []model.Post) error {
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
