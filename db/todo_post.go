package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ipreferwater/doko-graphql/model"
)

/*
TodoPostRepository is used to have a quick no-database test on GCP
I would like to test features but I'm currently stuck on connecting a GCP hosted database
to bypass this I created this dummy repo
*/
type TodoPostRepository struct {
	
}

var (
	posts []model.Post
)

func (r TodoPostRepository) GetPosts() ([]model.Post, error) {
	return posts, nil
}

func (r TodoPostRepository) CreatePosts(newPosts []*model.InputPost) error {
	for _, inputPost := range newPosts {

		fmt.Println("create a post")
		newPost := model.Post{
			ID:        len(posts),
			Title:     inputPost.Title,
			Text:      inputPost.Text,
			Latitude:  inputPost.Latitude,
			Longitude: inputPost.Longitude,
		}
		posts = append(posts, newPost)
		fmt.Printf("new len %d",len(posts))
	}
	return nil
}

func (n TodoPostRepository) UpdatePosts(postsToUpdate []model.Post) error {
	return nil
}
func (n TodoPostRepository) DeletePosts(idsPostToDelete []int) error {
	return nil
}

func (m TodoPostRepository) GetUserIdByUsernamePassword(userName string, password string) (int, error) {
	return 1, nil
}

func InitTodoPostRepository() {
	PostRepository = &TodoPostRepository{}
}
